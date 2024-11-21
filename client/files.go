package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
)

type UploadResultsRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	Data               []byte
	FileExtension      string
}

func (c *Client) UploadResults(ctx context.Context, req UploadResultsRequest) (*string, error) {
	res, err := c.productClient.UploadExtensionExecutionFile(ctx, &gen.UploadExtensionExecutionFileRequest{
		ProjectUuid:          req.ProjectUUID.String(),
		ProjectVersionUuid:   req.ProjectVersionUUID.String(),
		ExtensionUuid:        c.metadata.UUID,
		ExtensionVersionUuid: c.metadata.VersionUUID,
		ExecutionUuid:        req.ExecutionUUID.String(),
		Data:                 req.Data,
		FileExtension:        req.FileExtension,
	})

	if err != nil {
		return nil, err
	}

	return &res.Url, nil
}
