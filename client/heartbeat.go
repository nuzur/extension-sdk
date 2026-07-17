package client

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

const (
	// heartbeatInterval is how often an in-progress execution refreshes its
	// liveness. It must stay comfortably below the product reaper's
	// in-progress staleness window so a healthy run is never reclaimed.
	heartbeatInterval = 30 * time.Second
	// heartbeatTimeout bounds a single heartbeat RPC.
	heartbeatTimeout = 10 * time.Second
)

// StartHeartbeat periodically refreshes the execution's queue heartbeat until
// the returned stop function is called. Call it once an execution has been
// admitted (is INPROGRESS) and defer the stop for when the run finishes, so the
// queue reaper does not reclaim the slot while work is ongoing.
//
// authCtx must carry the caller's auth metadata as OUTGOING gRPC metadata and
// should outlive the run (e.g. a background context seeded from the incoming
// request's metadata), because the originating request context is cancelled as
// soon as the async StartExecution handler returns.
func (c *Client) StartHeartbeat(authCtx context.Context, executionUUID uuid.UUID) (stop func()) {
	ctx, cancel := context.WithCancel(authCtx)
	go func() {
		ticker := time.NewTicker(heartbeatInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				hbCtx, hbCancel := context.WithTimeout(ctx, heartbeatTimeout)
				_ = c.HeartbeatExecution(hbCtx, executionUUID)
				hbCancel()
			}
		}
	}()
	return cancel
}
