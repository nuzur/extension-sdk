package client

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
	nemgen "github.com/nuzur/extension-sdk/proto_deps/nem/idl/gen"
)

func (c *Client) GetExtension(ctx context.Context, extensionUUID uuid.UUID) (*nemgen.Extension, error) {
	return c.productClient.GetExtension(ctx, &gen.GetExtensionRequest{
		ExtensionUuid: extensionUUID.String(),
	})
}

func (c *Client) GetExtensionVersion(ctx context.Context, extensionVersionUUID uuid.UUID) (*nemgen.ExtensionVersion, error) {
	return c.productClient.GetExtensionVersion(ctx, &gen.GetExtensionVersionRequest{
		VersionUuid: extensionVersionUUID.String(),
	})
}

func (c *Client) GetExecution(ctx context.Context, extensionExecutionUUID uuid.UUID) (*nemgen.ExtensionExecution, error) {
	return c.productClient.GetExtensionExecution(ctx, &gen.GetExtensionExecutionRequest{
		ExecutionUuid: extensionExecutionUUID.String(),
	})
}

type CreateExecutionRequest struct {
	ProjectUUID          uuid.UUID
	ProjectVersionUUID   uuid.UUID
	ProjectExtensionUUID uuid.UUID
	Metadata             string
}

func (c *Client) CreateExecution(ctx context.Context, req CreateExecutionRequest) (*nemgen.ExtensionExecution, error) {
	if req.ProjectUUID == uuid.Nil {
		return nil, errors.New("project uuid is required")
	}

	if req.ProjectVersionUUID == uuid.Nil {
		return nil, errors.New("project version uuid is required")
	}
	return c.productClient.CreateExtensionExecution(ctx, &gen.CreateExtensionExecutionRequest{
		Execution: &nemgen.ExtensionExecution{
			ExtensionUuid:        c.metadata.UUID,
			ExtensionVersionUuid: c.metadata.VersionUUID,
			Status:               nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INPROGRESS,
			StatusMsg:            "in_progress",
			ProjectUuid:          req.ProjectUUID.String(),
			ProjectVersionUuid:   req.ProjectVersionUUID.String(),
			ProjectExtensionUuid: req.ProjectExtensionUUID.String(),
		},
	})
}

type UpdateExecutionRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	Status             nemgen.ExtensionExecutionStatus
	StatusMsg          string
	Metadata           string
}

func (c *Client) UpdateExecution(ctx context.Context, req UpdateExecutionRequest) (*nemgen.ExtensionExecution, error) {
	if req.ExecutionUUID == uuid.Nil {
		return nil, errors.New("execution uuid is required")
	}

	if req.ProjectUUID == uuid.Nil {
		return nil, errors.New("project uuid is required")
	}

	if req.ProjectVersionUUID == uuid.Nil {
		return nil, errors.New("project version uuid is required")
	}

	existingExec, err := c.productClient.GetExtensionExecution(ctx, &gen.GetExtensionExecutionRequest{
		ExecutionUuid: req.ExecutionUUID.String(),
	})

	if err != nil {
		return nil, err
	}

	if existingExec.ProjectUuid != req.ProjectUUID.String() || existingExec.ProjectVersionUuid != req.ProjectVersionUUID.String() {
		return nil, errors.New("project uuid missmatch")
	}

	return c.productClient.UpdateExtensionExecution(ctx, &gen.UpdateExtensionExecutionRequest{
		Execution: &nemgen.ExtensionExecution{
			Uuid:                 req.ExecutionUUID.String(),
			ExtensionUuid:        existingExec.ExtensionUuid,
			ExtensionVersionUuid: existingExec.ExtensionVersionUuid,
			ProjectUuid:          existingExec.ProjectUuid,
			ProjectVersionUuid:   existingExec.ProjectVersionUuid,
			ProjectExtensionUuid: existingExec.ProjectExtensionUuid,
			ExecutedByUuid:       existingExec.ExecutedByUuid,
			Metadata:             req.Metadata,
			Status:               req.Status,
			StatusMsg:            req.StatusMsg,
		},
	})

}
