package client

import (
	"context"
	"fmt"
	"path"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
	"github.com/nuzur/filetools"
)

type UploadResultsRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	Data               []byte
	FileExtension      string
}

func (c *Client) UploadExecutionResults(ctx context.Context, req UploadResultsRequest) (*string, error) {
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

	return &res.Filepath, nil
}

type DownloadExecutionResultsRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	FileExtension      string
}

type DownloadExecutionResultsResponse struct {
	FileDownloadUrl string
	LocalFilePath   string
}

func (c *Client) DownloadExecutionResults(ctx context.Context, req DownloadExecutionResultsRequest) (*DownloadExecutionResultsResponse, error) {
	// get the signed url from product
	res, err := c.productClient.GetExtensionExecutionFile(ctx, &gen.GetExtensionExecutionFileRequest{
		ProjectUuid:          req.ProjectUUID.String(),
		ProjectVersionUuid:   req.ProjectVersionUUID.String(),
		ExtensionUuid:        c.metadata.UUID,
		ExtensionVersionUuid: c.metadata.VersionUUID,
		ExecutionUuid:        req.ExecutionUUID.String(),
		FileExtension:        req.FileExtension,
	})
	if err != nil {
		return nil, err
	}

	// download into local file
	rootDir := filetools.CurrentLocalPath()
	filePath := path.Join(rootDir, "previous-executions", fmt.Sprintf("%s.%s", req.ExecutionUUID, req.FileExtension))

	err = filetools.DownloadFile(filePath, res.SignedUrl)
	if err != nil {
		return nil, err
	}

	return &DownloadExecutionResultsResponse{
		FileDownloadUrl: res.Filepath,
		LocalFilePath:   filePath,
	}, nil
}
