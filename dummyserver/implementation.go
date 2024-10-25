package dummyserver

import (
	"context"

	"github.com/gofrs/uuid"
	pb "github.com/nuzur/extension-sdk/proto/gen"
	"github.com/nuzur/extension-sdk/proto/nem/idl/gen"
)

// ActivateUserTeamInvite implements gen.NuzurProductServer.
func (s *Server) ActivateUserTeamInvite(context.Context, *pb.ActivateUserTeamInviteRequest) (*gen.UserTeam, error) {
	panic("unimplemented")
}

// CreateConnectionSecret implements gen.NuzurProductServer.
func (s *Server) CreateConnectionSecret(context.Context, *pb.CreateConnectionSecretRequest) (*gen.Connection, error) {
	panic("unimplemented")
}

// CreateDraftProjectVersionForExistingProject implements gen.NuzurProductServer.
func (s *Server) CreateDraftProjectVersionForExistingProject(context.Context, *pb.CreateDraftProjectVersionForExistingProjectRequest) (*gen.ProjectVersion, error) {
	panic("unimplemented")
}

// CreateProject implements gen.NuzurProductServer.
func (s *Server) CreateProject(context.Context, *pb.CreateProjectRequest) (*gen.Project, error) {
	panic("unimplemented")
}

// CreateProjectVersion implements gen.NuzurProductServer.
func (s *Server) CreateProjectVersion(context.Context, *pb.CreateProjectVersionRequest) (*gen.ProjectVersion, error) {
	panic("unimplemented")
}

// CreateTeam implements gen.NuzurProductServer.
func (s *Server) CreateTeam(context.Context, *pb.CreateTeamRequest) (*gen.Team, error) {
	panic("unimplemented")
}

// DeleteConnectionSecret implements gen.NuzurProductServer.
func (s *Server) DeleteConnectionSecret(context.Context, *pb.DeleteConnectionSecretRequest) (*gen.Connection, error) {
	panic("unimplemented")
}

// DiscardDraftProjectVersion implements gen.NuzurProductServer.
func (s *Server) DiscardDraftProjectVersion(context.Context, *pb.DiscardDraftProjectVersionRequest) (*gen.ProjectVersion, error) {
	panic("unimplemented")
}

// GetConnectionWithSecret implements gen.NuzurProductServer.
func (s *Server) GetConnectionWithSecret(context.Context, *pb.GetConnectionWithSecretRequest) (*gen.Connection, error) {
	panic("unimplemented")
}

// GetLatestProjectVersion implements gen.NuzurProductServer.
func (s *Server) GetLatestProjectVersion(context.Context, *pb.GetLatestProjectVersionRequest) (*gen.ProjectVersion, error) {
	panic("unimplemented")
}

// GetProjectForUser implements gen.NuzurProductServer.
func (s *Server) GetProjectForUser(context.Context, *pb.GetProjectForUserRequest) (*gen.Project, error) {
	return &gen.Project{
		Uuid: uuid.Must(uuid.NewV4()).String(),
		Name: "test",
	}, nil
}

// GetProjectVersionForUser implements gen.NuzurProductServer.
func (s *Server) GetProjectVersionForUser(context.Context, *pb.GetProjectVersionForUserRequest) (*gen.ProjectVersion, error) {
	return &gen.ProjectVersion{
		Uuid: uuid.Must(uuid.NewV4()).String(),
	}, nil
}

// GetTeamForUser implements gen.NuzurProductServer.
func (s *Server) GetTeamForUser(context.Context, *pb.GetTeamForUserRequest) (*gen.Team, error) {
	panic("unimplemented")
}

// GetUser implements gen.NuzurProductServer.
func (s *Server) GetUser(context.Context, *pb.GetUserRequest) (*gen.User, error) {
	panic("unimplemented")
}

// InviteUserToTeam implements gen.NuzurProductServer.
func (s *Server) InviteUserToTeam(context.Context, *pb.InviteUserToTeamRequest) (*gen.UserTeam, error) {
	panic("unimplemented")
}

// ListProjectVersions implements gen.NuzurProductServer.
func (s *Server) ListProjectVersions(context.Context, *pb.ListProjectVersionsRequest) (*pb.ListProjectVersionsResponse, error) {
	panic("unimplemented")
}

// ListProjectVersionsForUser implements gen.NuzurProductServer.
func (s *Server) ListProjectVersionsForUser(context.Context, *pb.ListProjectVersionsForUserRequest) (*pb.ListProjectVersionsForUserResponse, error) {
	panic("unimplemented")
}

// ListProjectsForUser implements gen.NuzurProductServer.
func (s *Server) ListProjectsForUser(context.Context, *pb.ListProjectsForUserRequest) (*pb.ListProjectsForUserResponse, error) {
	panic("unimplemented")
}

// ListTeamsForUser implements gen.NuzurProductServer.
func (s *Server) ListTeamsForUser(context.Context, *pb.ListTeamsForUserRequest) (*pb.ListTeamsForUserResponse, error) {
	panic("unimplemented")
}

// UpdateConnectionSecret implements gen.NuzurProductServer.
func (s *Server) UpdateConnectionSecret(context.Context, *pb.UpdateConnectionSecretRequest) (*gen.Connection, error) {
	panic("unimplemented")
}

// UpdateProject implements gen.NuzurProductServer.
func (s *Server) UpdateProject(context.Context, *pb.UpdateProjectRequest) (*gen.Project, error) {
	panic("unimplemented")
}

// UpdateProjectVersion implements gen.NuzurProductServer.
func (s *Server) UpdateProjectVersion(context.Context, *pb.UpdateProjectVersionRequest) (*gen.ProjectVersion, error) {
	panic("unimplemented")
}

// UpdateTeam implements gen.NuzurProductServer.
func (s *Server) UpdateTeam(context.Context, *pb.UpdateTeamRequest) (*gen.Team, error) {
	panic("unimplemented")
}

// UploadProjectSnapshot implements gen.NuzurProductServer.
func (s *Server) UploadProjectSnapshot(context.Context, *pb.UploadProjectSnapshotRequest) (*pb.UploadProjectSnapshotResponse, error) {
	panic("unimplemented")
}
