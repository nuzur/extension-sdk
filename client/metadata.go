package client

import (
	"context"

	pb "github.com/nuzur/extension-sdk/idl/gen"
)

func (c *Client) GetMetadata(context.Context, *pb.GetMetadataRequest) (*pb.GetMetadataResponse, error) {
	return &pb.GetMetadataResponse{
		Uuid:          c.metadata.UUID,
		VersionUuid:   c.metadata.VersionUUID,
		Identifier:    c.metadata.Identifier,
		DisplayName:   c.metadata.DisplayName,
		Description:   c.metadata.Description,
		NumberOfSteps: c.metadata.Steps,
	}, nil
}
