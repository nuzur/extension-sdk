package client

import (
	"context"
	"testing"

	"github.com/nuzur/extension-sdk/idl/gen"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestMetadata(t *testing.T) {
	client, server, err := BuildDummyClientServer()
	defer server.Stop()
	assert.NoError(t, err)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, AUTH_KEY, "some-key") // this should be passed by the called of the endpoint
	metadata, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{})

	assert.NoError(t, err)

	assert.NotNil(t, metadata)
	assert.NotNil(t, metadata.Uuid)
	assert.Equal(t, metadata.Name, "sdk")
	assert.Equal(t, metadata.Author, "nuzur")
	assert.Equal(t, metadata.NumberOfSteps, int32(1))
}
