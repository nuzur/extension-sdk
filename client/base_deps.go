package client

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	nemgen "github.com/nuzur/extension-sdk/proto_deps/nem/idl/gen"
	"golang.org/x/sync/errgroup"
)

type BaseDependenciesRequest struct {
	ProjectUUID        uuid.UUID
	ProjectVersionUUID uuid.UUID
}

type BaseDependenciesResponse struct {
	Extension        *nemgen.Extension
	ExtensionVersion *nemgen.ExtensionVersion
	Project          *nemgen.Project
	ProjectVersion   *nemgen.ProjectVersion
	Team             *nemgen.Team
}

func (c *Client) GetBaseDependencies(ctx context.Context, req BaseDependenciesRequest) (*BaseDependenciesResponse, error) {
	extensionUUID := uuid.FromStringOrNil(c.metadata.UUID)
	if extensionUUID == uuid.Nil {
		return nil, errors.New("invalid extension uuid")
	}
	extensionVersionUUID := uuid.FromStringOrNil(c.metadata.VersionUUID)
	if extensionVersionUUID == uuid.Nil {
		return nil, errors.New("invalid extension version uuid")
	}

	res := BaseDependenciesResponse{}
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		extension, err := c.GetExtension(ctx, extensionUUID)
		if err != nil {
			return err
		}
		res.Extension = extension
		return nil
	})

	eg.Go(func() error {
		extensionVersion, err := c.GetExtensionVersion(ctx, extensionVersionUUID)
		if err != nil {
			return err
		}
		res.ExtensionVersion = extensionVersion
		return nil
	})

	eg.Go(func() error {
		project, err := c.GetProject(ctx, req.ProjectUUID)
		if err != nil {
			return err
		}
		res.Project = project
		return nil
	})

	eg.Go(func() error {
		projectVersion, err := c.GetProjectVersion(ctx, req.ProjectVersionUUID)
		if err != nil {
			return err
		}
		res.ProjectVersion = projectVersion
		return nil
	})

	err := eg.Wait()

	if err != nil {
		return nil, err
	}

	if res.Project != nil {

		team, err := c.GetTeam(ctx, uuid.FromStringOrNil(res.Project.TeamUuid))
		if err != nil {
			return nil, err
		}
		res.Team = team
	}
	return &res, nil
}
