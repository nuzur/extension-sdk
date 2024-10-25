package client

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProject(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	project, err := client.GetProject(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, project)
	assert.NotNil(t, project.Uuid)
	assert.Equal(t, project.Name, "test")
}

func TestProjectVersion(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	projectVersion, err := client.GetProjectVersion(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, projectVersion)
	assert.NotNil(t, projectVersion.Uuid)
}
