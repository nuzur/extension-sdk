package client

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	pb "github.com/nuzur/extension-sdk/idl/gen"
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
	exec, err := c.productClient.GetExtensionExecution(ctx, &gen.GetExtensionExecutionRequest{
		ExecutionUuid: extensionExecutionUUID.String(),
	})

	if err != nil {
		return nil, err
	}

	if exec.ExtensionUuid != c.metadata.UUID {
		return nil, errors.New("extension uuid missmatch")
	}

	return exec, nil
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
			Metadata:             req.Metadata,
		},
	})
}

type UpdateExecutionRequest struct {
	ExecutionUUID      uuid.UUID
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
	Status             pb.ExecutionStatus
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
			Status:               mapExtensionStatus(req.Status),
			StatusMsg:            req.StatusMsg,
		},
	})

}

func mapExtensionStatus(in pb.ExecutionStatus) nemgen.ExtensionExecutionStatus {
	switch in {
	case pb.ExecutionStatus_EXECUTION_STATUS_INVALID:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INVALID
	case pb.ExecutionStatus_EXECUTION_STATUS_INPROGRESS:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INPROGRESS
	case pb.ExecutionStatus_EXECUTION_STATUS_SUCCEEDED:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_SUCCEEDED
	case pb.ExecutionStatus_EXECUTION_STATUS_FAILED:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_FAILED
	case pb.ExecutionStatus_EXECUTION_STATUS_CANCELLED:
		return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_CANCELLED
	}
	return nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INVALID
}
