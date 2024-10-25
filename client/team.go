package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
	nemgen "github.com/nuzur/extension-sdk/proto_deps/nem/idl/gen"
	"google.golang.org/grpc/metadata"
)

func (c *Client) GetTeam(ctx context.Context, teamUUID uuid.UUID) (*nemgen.Team, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, c.token)
	return c.productClient.GetTeamForUser(ctx, &gen.GetTeamForUserRequest{
		TeamUuid: teamUUID.String(),
	})
}
