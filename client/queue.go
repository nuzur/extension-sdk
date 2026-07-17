package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	pb "github.com/nuzur/extension-sdk/idl/gen"
	nemgen "github.com/nuzur/nem/idl/gen"
	"google.golang.org/grpc/metadata"
)

// RunFunc executes an admitted execution to completion. Each extension server
// provides one wrapping its manager's run entrypoint, so the shared queue
// machinery stays extension-agnostic. A non-nil error is recorded as a failed
// execution.
type RunFunc func(ctx context.Context, req *pb.StartExecutionRequest, exec *nemgen.ExtensionExecution) error

// QueueCoordinator drives the client side of the global generation admission
// queue for an extension server: it admits or queues freshly-created
// executions, starts them once a slot is granted, keeps their heartbeats fresh
// while running, and tracks in-flight runs for graceful drain. The queue policy
// (cap, priority, reaping) lives in product; this coordinates against it.
//
// It is shared by every extension server so the queueing behaviour is written
// once. The only per-extension difference — how a run is actually executed — is
// injected as a RunFunc.
type QueueCoordinator struct {
	client *Client
	run    RunFunc

	mu       sync.Mutex
	pending  map[string]*pb.StartExecutionRequest
	spawned  map[string]bool
	inflight sync.WaitGroup
}

func NewQueueCoordinator(c *Client, run RunFunc) *QueueCoordinator {
	return &QueueCoordinator{
		client:  c,
		run:     run,
		pending: make(map[string]*pb.StartExecutionRequest),
		spawned: make(map[string]bool),
	}
}

// AdmitOrQueue is called from StartExecution once the QUEUED execution row
// exists. On admission it starts the run and returns an async "starting"
// response; otherwise it holds the request for a later poll and returns an async
// "queued" response (carrying the queue position) the client keeps polling on.
func (q *QueueCoordinator) AdmitOrQueue(ctx context.Context, req *pb.StartExecutionRequest, exec *nemgen.ExtensionExecution) (*pb.StartExecutionResponse, error) {
	admission, err := q.client.AcquireExecutionSlot(ctx, uuid.FromStringOrNil(exec.Uuid))
	if err != nil {
		return nil, err
	}

	if admission.State == SlotAdmitted {
		q.ensureRun(ctx, req, exec)
		return &pb.StartExecutionResponse{
			ExecutionUuid: exec.Uuid,
			Type:          pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_ASYNC,
			Data: &pb.ExecutionResponseTypeData{
				Async: &pb.ExecutionResponseTypeAsyncData{
					StatusMessage: q.client.Localize("resolving_config_values", req.Locale, "Resolving Config Values"),
				},
			},
		}, nil
	}

	q.mu.Lock()
	q.pending[exec.Uuid] = req
	q.mu.Unlock()
	return &pb.StartExecutionResponse{
		ExecutionUuid: exec.Uuid,
		Type:          pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_ASYNC,
		Data: &pb.ExecutionResponseTypeData{
			Async: queuedAsyncData(q.client, req.Locale, admission),
		},
	}, nil
}

// AdmitOnPoll is called from GetExecution for an execution this process is not
// yet running. It re-evaluates admission (also refreshing the queue heartbeat);
// on admission it starts the held request, otherwise it reports the current
// queue position. Returns a terminal failure if the entry is no longer
// queueable.
func (q *QueueCoordinator) AdmitOnPoll(ctx context.Context, executionUUID string) (*pb.GetExecutionResponse, error) {
	q.mu.Lock()
	req, held := q.pending[executionUUID]
	q.mu.Unlock()

	admission, err := q.client.AcquireExecutionSlot(ctx, uuid.FromStringOrNil(executionUUID))
	if err != nil {
		return finalFailure(q.client.Localize("execution_expired", localeOf(req), "Execution is no longer available")), nil
	}

	if admission.State == SlotAdmitted {
		if !held {
			// Admitted but we never held the request (e.g. this process restarted
			// while it was queued); it cannot be reconstructed here.
			return finalFailure(q.client.Localize("execution_lost", "", "Execution could not be resumed; please retry")), nil
		}
		exec, err := q.client.GetExecution(ctx, uuid.FromStringOrNil(executionUUID))
		if err != nil {
			return nil, err
		}
		q.ensureRun(ctx, req, exec)
		q.mu.Lock()
		delete(q.pending, executionUUID)
		q.mu.Unlock()
		msg := q.client.Localize("resolving_config_values", req.Locale, "Resolving Config Values")
		return &pb.GetExecutionResponse{
			Status:    pb.ExecutionStatus_EXECUTION_STATUS_INPROGRESS,
			StatusMsg: msg,
			Type:      pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_ASYNC,
			Data: &pb.ExecutionResponseTypeData{
				Async: &pb.ExecutionResponseTypeAsyncData{StatusMessage: msg},
			},
		}, nil
	}

	// Still queued. Report as an in-progress-style async response so older
	// clients keep polling, while the structured queue fields let newer clients
	// render a dedicated waiting state.
	locale := localeOf(req)
	return &pb.GetExecutionResponse{
		Status:    pb.ExecutionStatus_EXECUTION_STATUS_INPROGRESS,
		StatusMsg: queuedStatusMessage(q.client, locale, admission),
		Type:      pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_ASYNC,
		Data: &pb.ExecutionResponseTypeData{
			Async: queuedAsyncData(q.client, locale, admission),
		},
	}, nil
}

// ensureRun starts the execution's run goroutine exactly once, keeping its
// heartbeat fresh for the duration.
func (q *QueueCoordinator) ensureRun(ctx context.Context, req *pb.StartExecutionRequest, exec *nemgen.ExtensionExecution) {
	q.mu.Lock()
	if q.spawned[exec.Uuid] {
		q.mu.Unlock()
		return
	}
	q.spawned[exec.Uuid] = true
	q.mu.Unlock()

	// Detach the auth metadata onto a background context so both the heartbeat
	// and the run outlive this request.
	md, _ := metadata.FromIncomingContext(ctx)
	bctx := metadata.NewOutgoingContext(context.Background(), md)

	q.inflight.Add(1)
	go func() {
		defer q.inflight.Done()

		stopHeartbeat := q.client.StartHeartbeat(bctx, uuid.FromStringOrNil(exec.Uuid))
		defer stopHeartbeat()

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("failed to run: %v\n", r)
				q.markFailed(bctx, req, exec, fmt.Sprintf("Failed to run (panic): %v", r))
			}
		}()

		if err := q.run(ctx, req, exec); err != nil {
			q.markFailed(bctx, req, exec, "Failed to start: "+err.Error())
		}
	}()
}

func (q *QueueCoordinator) markFailed(bctx context.Context, req *pb.StartExecutionRequest, exec *nemgen.ExtensionExecution, msg string) {
	_, _ = q.client.UpdateExecution(bctx, UpdateExecutionRequest{
		ExecutionUUID:      uuid.FromStringOrNil(exec.Uuid),
		ProjectUUID:        uuid.FromStringOrNil(req.ProjectUuid),
		ProjectVersionUUID: uuid.FromStringOrNil(req.ProjectVersionUuid),
		Status:             pb.ExecutionStatus_EXECUTION_STATUS_FAILED,
		StatusMsg:          msg,
	})
}

// Drain waits for all in-flight runs to finish, or until ctx is done. Extension
// servers call this from their fx OnStop hook.
func (q *QueueCoordinator) Drain(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		q.inflight.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-ctx.Done():
	}
}

func localeOf(req *pb.StartExecutionRequest) string {
	if req != nil {
		return req.Locale
	}
	return ""
}

func queuedStatusMessage(c *Client, locale string, admission *AcquireSlotResponse) string {
	def := fmt.Sprintf("Queued — position %d", admission.Position)
	if admission.EstimatedWaitSeconds > 0 {
		def = fmt.Sprintf("Queued — position %d (about %ds)", admission.Position, admission.EstimatedWaitSeconds)
	}
	return c.Localize("execution_queued", locale, def)
}

func queuedAsyncData(c *Client, locale string, admission *AcquireSlotResponse) *pb.ExecutionResponseTypeAsyncData {
	return &pb.ExecutionResponseTypeAsyncData{
		StatusMessage:        queuedStatusMessage(c, locale, admission),
		Queued:               true,
		QueuePosition:        admission.Position,
		EstimatedWaitSeconds: admission.EstimatedWaitSeconds,
	}
}

func finalFailure(msg string) *pb.GetExecutionResponse {
	return &pb.GetExecutionResponse{
		Status:    pb.ExecutionStatus_EXECUTION_STATUS_FAILED,
		StatusMsg: msg,
		Type:      pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_FINAL,
		Data: &pb.ExecutionResponseTypeData{
			Final: &pb.ExecutionResponseTypeFinalData{
				Status:        pb.ExecutionStatus_EXECUTION_STATUS_FAILED,
				StatusMessage: msg,
			},
		},
	}
}
