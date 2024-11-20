package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/v1/proto_deps/gen"
	nemgen "github.com/nuzur/extension-sdk/v1/proto_deps/nem/idl/gen"
)

func (c *Client) GetTeam(ctx context.Context, teamUUID uuid.UUID) (*nemgen.Team, error) {
	return c.productClient.GetTeamForUser(ctx, &gen.GetTeamForUserRequest{
		TeamUuid: teamUUID.String(),
	})
}
