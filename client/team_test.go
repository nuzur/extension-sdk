package client

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTeam(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	team, err := client.GetTeam(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, team)
	assert.NotNil(t, team.Uuid)
}
