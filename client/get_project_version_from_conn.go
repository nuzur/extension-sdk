package client

import (
	"context"

	"github.com/nuzur/extension-sdk/proto_deps/gen"
	nemgen "github.com/nuzur/nem/idl/gen"
)

func (c *Client) GetProjectVersionFromConnection(ctx context.Context, req *gen.GetProjectVersionFromConnectionRequest) (*nemgen.ProjectVersion, error) {
	res, err := c.connectionManagerClient.GetProjectVersionFromConnection(
		ctx,
		req,
	)

	if err != nil {
		return nil, err
	}

	return res.ProjectVersion, nil
}
