package dummyserver

import (
	"context"

	"github.com/gofrs/uuid"
	pb "github.com/nuzur/extension-sdk/v1/proto_deps/gen"
	"github.com/nuzur/extension-sdk/v1/proto_deps/nem/idl/gen"
)

const DummyProjectUUID = "51ffef0c-0b21-4bed-bd3e-c17c2314a19b"
const DummyProjectVersionUUID = "9605d1ac-9a44-4642-aa68-0a6d92650b91"

// GetProjectForUser implements gen.NuzurProductServer.
func (s *Server) GetProjectForUser(ctx context.Context, req *pb.GetProjectForUserRequest) (*gen.Project, error) {
	return &gen.Project{
		Uuid: req.ProjectUuid,
		Name: "test",
	}, nil
}

// GetProjectVersionForUser implements gen.NuzurProductServer.
func (s *Server) GetProjectVersionForUser(ctx context.Context, req *pb.GetProjectVersionForUserRequest) (*gen.ProjectVersion, error) {
	return &gen.ProjectVersion{
		Uuid: req.ProjectVersionUuid,
	}, nil
}

// GetTeamForUser implements gen.NuzurProductServer.
func (s *Server) GetTeamForUser(ctx context.Context, req *pb.GetTeamForUserRequest) (*gen.Team, error) {
	return &gen.Team{
		Uuid: req.TeamUuid,
	}, nil
}

func (s *Server) GetExtension(ctx context.Context, req *pb.GetExtensionRequest) (*gen.Extension, error) {
	return &gen.Extension{
		Uuid:       req.ExtensionUuid,
		Identifier: "test_extension",
	}, nil
}

func (s *Server) GetExtensionVersion(ctx context.Context, req *pb.GetExtensionVersionRequest) (*gen.ExtensionVersion, error) {
	return &gen.ExtensionVersion{
		Uuid: req.VersionUuid,
	}, nil
}

func (s *Server) GetExtensionExecution(ctx context.Context, req *pb.GetExtensionExecutionRequest) (*gen.ExtensionExecution, error) {
	return &gen.ExtensionExecution{
		Uuid:               req.ExecutionUuid,
		ProjectUuid:        uuid.FromStringOrNil(DummyProjectUUID).String(),
		ProjectVersionUuid: uuid.FromStringOrNil(DummyProjectVersionUUID).String(),
		ExtensionUuid:      "0e4a4d9f-dd90-4173-b737-eeb75395a61a",
	}, nil
}
func (s *Server) CreateExtensionExecution(ctx context.Context, req *pb.CreateExtensionExecutionRequest) (*gen.ExtensionExecution, error) {
	return req.Execution, nil
}
func (s *Server) UpdateExtensionExecution(ctx context.Context, req *pb.UpdateExtensionExecutionRequest) (*gen.ExtensionExecution, error) {
	return req.Execution, nil
}
