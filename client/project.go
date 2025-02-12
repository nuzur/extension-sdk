package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
	nemgen "github.com/nuzur/nem/idl/gen"
)

func (c *Client) GetProject(ctx context.Context, projectUUID uuid.UUID) (*nemgen.Project, error) {
	return c.productClient.GetProjectForUser(ctx, &gen.GetProjectForUserRequest{
		ProjectUuid: projectUUID.String(),
	})
}

func (c *Client) GetProjectVersion(ctx context.Context, projectVersionUUID uuid.UUID) (*nemgen.ProjectVersion, error) {
	return c.productClient.GetProjectVersionForUser(ctx, &gen.GetProjectVersionForUserRequest{
		ProjectVersionUuid: projectVersionUUID.String(),
	})
}
