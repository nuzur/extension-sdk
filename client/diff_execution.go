package client

import (
	"context"

	"github.com/nuzur/extension-sdk/proto_deps/gen"
)

func (c *Client) StartChangesDiff(ctx context.Context, req *gen.StartChangesDiffRequest) (*gen.StartChangesDiffResponse, error) {
	return c.connectionManagerClient.StartChangesDiff(ctx, req)
}

func (c *Client) GetChangesDiff(ctx context.Context, req *gen.GetChangesDiffRequest) (*gen.GetChangesDiffResponse, error) {
	return c.connectionManagerClient.GetChangesDiff(ctx, req)
}

func (c *Client) CancelChangesDiff(ctx context.Context, req *gen.CancelChangesDiffRequest) (*gen.CancelChangesDiffResponse, error) {
	return c.connectionManagerClient.CancelChangesDiff(ctx, req)
}
