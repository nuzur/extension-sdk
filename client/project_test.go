package client

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestProject(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
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
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
	projectVersion, err := client.GetProjectVersion(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, projectVersion)
	assert.NotNil(t, projectVersion.Uuid)
}
