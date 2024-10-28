package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/nem/idl/gen"
)

type ResolveConfigValuesRequest struct {
	ProjectUUID          uuid.UUID
	ProjectExtensionUUID uuid.UUID
	RawConfigValues      map[string]string
}

func (c *Client) ResolveConfigValues(ctx context.Context, req ResolveConfigValuesRequest) (map[string]string, error) {
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

		res := make(map[string]string)
		err = json.Unmarshal([]byte(projectExtension.ConfigurationEntityValues), &res)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return req.RawConfigValues, nil
}
