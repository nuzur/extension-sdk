package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/v1/proto_deps/nem/idl/gen"
)

type ResolveConfigValuesRequest struct {
	ProjectUUID          uuid.UUID
	ProjectExtensionUUID uuid.UUID
	RawConfigValues      string
}

func (c *Client) ResolveConfigValues(ctx context.Context, req ResolveConfigValuesRequest, values any) error {
	if req.ProjectExtensionUUID == uuid.Nil && len(req.RawConfigValues) == 0 {
		return errors.New("missing config data")
	}

	if req.ProjectExtensionUUID != uuid.Nil {
		project, err := c.GetProject(ctx, req.ProjectUUID)
		if err != nil {
			return err
		}

		var projectExtension *gen.ProjectExtension = nil
		for _, pe := range project.ProjectExtensions {
			if pe.Uuid == req.ProjectExtensionUUID.String() {
				projectExtension = pe
			}
		}

		if projectExtension == nil {
			return errors.New("project extension not found in project")
		}

		err = json.Unmarshal([]byte(projectExtension.ConfigurationEntityValues), &values)
		if err != nil {
			return err
		}

		return nil
	}

	err := json.Unmarshal([]byte(req.RawConfigValues), &values)
	if err != nil {
		return err
	}

	return nil
}
