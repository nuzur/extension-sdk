package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto/gen"
	nemgen "github.com/nuzur/extension-sdk/proto/nem/idl/gen"
	"google.golang.org/grpc/metadata"
)

func (c *Client) GetProject(ctx context.Context, projectUUID uuid.UUID) (*nemgen.Project, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, c.token)
	return c.productClient.GetProjectForUser(ctx, &gen.GetProjectForUserRequest{
		ProjectUuid: projectUUID.String(),
	})
}

func (c *Client) GetProjectVersion(ctx context.Context, projectVersionUUID uuid.UUID) (*nemgen.ProjectVersion, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, c.token)
	return c.productClient.GetProjectVersionForUser(ctx, &gen.GetProjectVersionForUserRequest{
		ProjectVersionUuid: projectVersionUUID.String(),
	})
}
