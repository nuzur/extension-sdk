package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto/gen"
	nemgen "github.com/nuzur/extension-sdk/proto/nem/idl/gen"
	"google.golang.org/grpc/metadata"
)

func (c *Client) GetProject(ctx context.Context, projectUUID uuid.UUID) (*nemgen.Project, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", c.token)
	return c.productClient.GetProjectForUser(ctx, &gen.GetProjectForUserRequest{
		ProjectUuid: projectUUID.String(),
	})
}
