package client

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	pb "github.com/nuzur/extension-sdk/idl/gen"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
	nemgen "github.com/nuzur/nem/idl/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Client) GetExtension(ctx context.Context, extensionUUID uuid.UUID) (*nemgen.Extension, error) {
	return c.productClient.GetExtension(ctx, &gen.GetExtensionRequest{
		ExtensionUuid: extensionUUID.String(),
	})
}

func (c *Client) GetExtensionVersion(ctx context.Context, extensionVersionUUID uuid.UUID) (*nemgen.ExtensionVersion, error) {
	return c.productClient.GetExtensionVersion(ctx, &gen.GetExtensionVersionRequest{
		VersionUuid: extensionVersionUUID.String(),
	})
}

func (c *Client) GetExecution(ctx context.Context, extensionExecutionUUID uuid.UUID) (*nemgen.ExtensionExecution, error) {
	exec, err := c.productClient.GetExtensionExecution(ctx, &gen.GetExtensionExecutionRequest{
		ExecutionUuid: extensionExecutionUUID.String(),
	})

	if err != nil {
		return nil, err
	}

	if exec.ExtensionUuid != c.metadata.UUID {
		return nil, errors.New("extension uuid missmatch")
	}

	return exec, nil
}

type CreateExecutionRequest struct {
	ProjectUUID          uuid.UUID
	ProjectVersionUUID   uuid.UUID
	ProjectExtensionUUID uuid.UUID
	Metadata             string
	Extension            *nemgen.Extension
}

// CreateExecution creates the execution row in the QUEUED state. Admission into
// the global generation queue (and promotion to INPROGRESS) is handled
// separately via AcquireExecutionSlot, so the extension server does not begin
// work until it is granted a slot. The former monthly Pro-execution limit has
// been replaced by the queue.
func (c *Client) CreateExecution(ctx context.Context, req CreateExecutionRequest) (*nemgen.ExtensionExecution, error) {
	if req.ProjectUUID == uuid.Nil {
		return nil, errors.New("project uuid is required")
	}

	if req.ProjectVersionUUID == uuid.Nil {
		return nil, errors.New("project version uuid is required")
	}

	return c.productClient.CreateExtensionExecution(ctx, &gen.CreateExtensionExecutionRequest{
		Execution: &nemgen.ExtensionExecution{
			ExtensionUuid:        c.metadata.UUID,
			ExtensionVersionUuid: c.metadata.VersionUUID,
			Status:               nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_QUEUED,
			StatusMsg:            "queued",
			ProjectUuid:          req.ProjectUUID.String(),
			ProjectVersionUuid:   req.ProjectVersionUUID.String(),
			ProjectExtensionUuid: req.ProjectExtensionUUID.String(),
			Metadata:             req.Metadata,
			// Seed the heartbeat at creation so the freshly-queued row is
			// immediately "live" and not mistaken for abandoned before the first
			// AcquireExecutionSlot poll refreshes it.
			LastHeartbeatAt: timestamppb.Now(),
		},
	})
}

// SlotState mirrors the product admission-queue slot state.
type SlotState int

const (
	SlotInvalid SlotState = iota
	SlotAdmitted
	SlotQueued
)

// AcquireSlotResponse is the outcome of an AcquireExecutionSlot call.
type AcquireSlotResponse struct {
	State                SlotState
	Position             int64
	EstimatedWaitSeconds int64
}

// AcquireExecutionSlot asks product whether this execution may start now. It is
// safe to call repeatedly while waiting: each call refreshes the queue
// heartbeat and re-evaluates admission. When State is SlotAdmitted the row has
// been promoted to INPROGRESS and work may begin.
func (c *Client) AcquireExecutionSlot(ctx context.Context, executionUUID uuid.UUID) (*AcquireSlotResponse, error) {
	if executionUUID == uuid.Nil {
		return nil, errors.New("execution uuid is required")
	}
	res, err := c.productClient.AcquireExecutionSlot(ctx, &gen.AcquireExecutionSlotRequest{
		ExecutionUuid: executionUUID.String(),
	})
	if err != nil {
		return nil, err
	}
	var state SlotState
	switch res.State {
	case gen.ExecutionSlotState_EXECUTION_SLOT_STATE_ADMITTED:
		state = SlotAdmitted
	case gen.ExecutionSlotState_EXECUTION_SLOT_STATE_QUEUED:
		state = SlotQueued
	default:
		state = SlotInvalid
	}
	return &AcquireSlotResponse{
		State:                state,
		Position:             res.Position,
		EstimatedWaitSeconds: res.EstimatedWaitSeconds,
	}, nil
}

// HeartbeatExecution refreshes the liveness timestamp of an in-progress
// execution so the queue reaper does not reclaim its slot mid-run.
func (c *Client) HeartbeatExecution(ctx context.Context, executionUUID uuid.UUID) error {
	if executionUUID == uuid.Nil {
		return errors.New("execution uuid is required")
	}
	_, err := c.productClient.HeartbeatExecution(ctx, &gen.HeartbeatExecutionRequest{
		ExecutionUuid: executionUUID.String(),
	})
	return err
}

type CheckLimitRequest struct {
	ProjectUUID   uuid.UUID
	ExtensionUUID uuid.UUID
}

type CheckLimitResponse struct {
	IsLimited bool
	Limit     int64
	Current   int64
}

func (c *Client) CheckExtensionExecutionLimit(ctx context.Context, req CheckLimitRequest) (*CheckLimitResponse, error) {
	if req.ProjectUUID == uuid.Nil {
		return nil, errors.New("project uuid is required")
	}

	extUUID := c.metadata.UUID
	if req.ExtensionUUID != uuid.Nil {
		extUUID = req.ExtensionUUID.String()
	}

	res, err := c.productClient.CheckExtensionExecutionLimit(ctx, &gen.CheckExtensionExecutionLimitRequest{
		ProjectUuid:   req.ProjectUUID.String(),
		ExtensionUuid: extUUID,
	})
	if err != nil {
		return nil, err
	}

	return &CheckLimitResponse{
		IsLimited: res.IsLimited,
		Limit:     res.Limit,
		Current:   res.Current,
	}, nil
}

type UpdateExecutionRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	Status             pb.ExecutionStatus
	StatusMsg          string
	Metadata           string
}

func (c *Client) UpdateExecution(ctx context.Context, req UpdateExecutionRequest) (*nemgen.ExtensionExecution, error) {
	if req.ExecutionUUID == uuid.Nil {
		return nil, errors.New("execution uuid is required")
	}

	if req.ProjectUUID == uuid.Nil {
		return nil, errors.New("project uuid is required")
	}

	if req.ProjectVersionUUID == uuid.Nil {
		return nil, errors.New("project version uuid is required")
	}

	existingExec, err := c.productClient.GetExtensionExecution(ctx, &gen.GetExtensionExecutionRequest{
		ExecutionUuid: req.ExecutionUUID.String(),
	})

	if err != nil {
		return nil, err
	}

	if existingExec.ProjectUuid != req.ProjectUUID.String() || existingExec.ProjectVersionUuid != req.ProjectVersionUUID.String() {
		return nil, errors.New("project uuid missmatch")
	}

	return c.productClient.UpdateExtensionExecution(ctx, &gen.UpdateExtensionExecutionRequest{
		Execution: &nemgen.ExtensionExecution{
			Uuid:                 req.ExecutionUUID.String(),
			ExtensionUuid:        existingExec.ExtensionUuid,
			ExtensionVersionUuid: existingExec.ExtensionVersionUuid,
			ProjectUuid:          existingExec.ProjectUuid,
			ProjectVersionUuid:   existingExec.ProjectVersionUuid,
			ProjectExtensionUuid: existingExec.ProjectExtensionUuid,
			ExecutedByUuid:       existingExec.ExecutedByUuid,
			Metadata:             req.Metadata,
			Status:               mapExtensionStatus(req.Status),
			StatusMsg:            req.StatusMsg,
			UpdatedAt:            timestamppb.Now(),
		},
	})

}

func mapExtensionStatus(in pb.ExecutionStatus) nemgen.ExtensionExecutionStatus {
	switch in {
	case pb.ExecutionStatus_EXECUTION_STATUS_INVALID:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INVALID
	case pb.ExecutionStatus_EXECUTION_STATUS_INPROGRESS:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INPROGRESS
	case pb.ExecutionStatus_EXECUTION_STATUS_SUCCEEDED:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_SUCCEEDED
	case pb.ExecutionStatus_EXECUTION_STATUS_FAILED:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_FAILED
	case pb.ExecutionStatus_EXECUTION_STATUS_CANCELLED:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_CANCELLED
	}
	return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INVALID
}
