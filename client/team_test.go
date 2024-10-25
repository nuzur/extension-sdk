package client

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestTeam(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
	team, err := client.GetTeam(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, team)
	assert.NotNil(t, team.Uuid)
}
