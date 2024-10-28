package client

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/nem/idl/gen"
)

type ResolveConfigValuesRequest struct {
	ProjectUUID          uuid.UUID
	ProjectExtensionUUID uuid.UUID
	RawConfigValues      string
}

func (c *Client) ResolveConfigValues(ctx context.Context, req ResolveConfigValuesRequest) (*string, error) {
	if req.ProjectExtensionUUID == uuid.Nil && len(req.RawConfigValues) == 0 {
		return nil, errors.New("missing config data")
	}

	if req.ProjectExtensionUUID != uuid.Nil {
		project, err := c.GetProject(ctx, req.ProjectUUID)
		if err != nil {
			return nil, err
		}

		var projectExtension *gen.ProjectExtension = nil
		for _, pe := range project.ProjectExtensions {
			if pe.Uuid == req.ProjectExtensionUUID.String() {
				projectExtension = pe
			}
		}

		if projectExtension == nil {
			return nil, errors.New("project extension not found in project")
		}

		return &projectExtension.ConfigurationEntityValues, nil
	}

	return &req.RawConfigValues, nil
}
