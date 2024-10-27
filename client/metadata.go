package client

import (
	"context"

	pb "github.com/nuzur/extension-sdk/idl/gen"
)

func (c *Client) GetMetadata(context.Context, *pb.GetMetadataRequest) (*pb.GetMetadataResponse, error) {
	return &pb.GetMetadataResponse{
		Uuid:          c.metadata.UUID,
		VersionUuid:   c.metadata.VersionUUID,
		Name:          c.metadata.Name,
		Author:        c.metadata.Author,
		NumberOfSteps: c.metadata.Steps,
	}, nil
}
