package client

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/v1/dummyserver"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestExtension(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
	extension, err := client.GetExtension(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, extension)
	assert.NotNil(t, extension.Uuid)
	assert.Equal(t, extension.Identifier, "test_extension")
}

func TestExtensionVersion(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
	extensionVersion, err := client.GetExtensionVersion(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, extensionVersion)
	assert.NotNil(t, extensionVersion.Uuid)
}

func TestGetExecution(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
	exec, err := client.GetExecution(ctx, uuid.Must(uuid.NewV4()))

	assert.NoError(t, err)

	assert.NotNil(t, exec)
	assert.NotNil(t, exec.Uuid)
}

func TestCreateExecution(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint

	exec, err := client.CreateExecution(ctx, CreateExecutionRequest{})
	assert.Error(t, err)
	assert.Nil(t, exec)

	exec, err = client.CreateExecution(ctx, CreateExecutionRequest{
		ProjectUUID: uuid.Must(uuid.NewV4()),
	})
	assert.Error(t, err)
	assert.Nil(t, exec)

	exec, err = client.CreateExecution(ctx, CreateExecutionRequest{
		ProjectUUID:        uuid.Must(uuid.NewV4()),
		ProjectVersionUUID: uuid.Must(uuid.NewV4()),
	})
	assert.NoError(t, err)

	assert.NotNil(t, exec)
	assert.NotNil(t, exec.Uuid)
}

func TestUpdateExecution(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint

	exec, err := client.UpdateExecution(ctx, UpdateExecutionRequest{})
	assert.Error(t, err)
	assert.Nil(t, exec)

	exec, err = client.UpdateExecution(ctx, UpdateExecutionRequest{
		ExecutionUUID: uuid.Must(uuid.NewV4()),
	})
	assert.Error(t, err)
	assert.Nil(t, exec)

	exec, err = client.UpdateExecution(ctx, UpdateExecutionRequest{
		ExecutionUUID:      uuid.Must(uuid.NewV4()),
		ProjectUUID:        uuid.FromStringOrNil(dummyserver.DummyProjectUUID),
		ProjectVersionUUID: uuid.FromStringOrNil(dummyserver.DummyProjectVersionUUID),
	})
	assert.NoError(t, err)

	assert.NotNil(t, exec)
	assert.NotNil(t, exec.Uuid)
}
