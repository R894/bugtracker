package grpc

import (
	"bugtracker/internal/domain/bug/core/domain/aggregate"
	"bugtracker/internal/domain/bug/ports"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *server) SaveBug(ctx context.Context, in *ports.SaveBugRequest) (*ports.EmptyResponse, error) {
	bug := aggregate.NewBug(aggregate.CreateBugRequest{Title: in.Title, Description: in.Description, ProjectId: in.ProjectId})
	s.d.SaveBug(ctx, bug)
	return &ports.EmptyResponse{}, nil
}

func (s *server) GetBugById(ctx context.Context, in *ports.GetBugByIDRequest) (*ports.BugResponse, error) {
	bug, err := s.d.GetBugByID(ctx, in.BugId)
	if err != nil {
		return nil, err
	}

	return &ports.BugResponse{BugId: bug.ID, Title: bug.Title, Description: bug.Description, Status: string(bug.Status), Priority: string(bug.Priority), Assignee: bug.Assignee, ProjectId: bug.ProjectId, CreatedAt: timestamppb.New(bug.CreatedAt), UpdatedAt: timestamppb.New(bug.UpdatedAt)}, nil
}

func (s *server) GetBugByProjectId(ctx context.Context, in *ports.GetBugsByProjectIDRequest) (*ports.GetBugsByProjectIDResponse, error) {
	bugs, err := s.d.GetBugsByProjectID(ctx, in.ProjectId)
	if err != nil {
		return nil, err
	}

	var convertedBugs []*ports.BugResponse
	for _, bug := range bugs {
		convertedBugs = append(convertedBugs, &ports.BugResponse{BugId: bug.ID, Title: bug.Title, Description: bug.Description, Status: string(bug.Status), Priority: string(bug.Priority), Assignee: bug.Assignee, ProjectId: bug.ProjectId, CreatedAt: timestamppb.New(bug.CreatedAt), UpdatedAt: timestamppb.New(bug.UpdatedAt)})
	}
	return &ports.GetBugsByProjectIDResponse{Bugs: convertedBugs}, nil
}

func (s *server) AssignBugTo(ctx context.Context, in *ports.AssignBugToRequest) (*ports.EmptyResponse, error) {
	err := s.d.AssignBugTo(ctx, in.BugId, in.Username)
	if err != nil {
		return nil, err
	}
	return &ports.EmptyResponse{}, nil
}
