package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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

	return &res.Url, nil
}

type GetExecutionResultsRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	FileExtension      string
}

func (c *Client) DownloadExecutionResults(ctx context.Context, req GetExecutionResultsRequest) (*string, error) {
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
	filePath := path.Join("previous-executions", fmt.Sprintf("%s.%s", req.ExecutionUUID, req.FileExtension))

	// create the file
	out, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	// download
	resp, err := http.Get(res.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// write to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}

	return &filePath, nil
}