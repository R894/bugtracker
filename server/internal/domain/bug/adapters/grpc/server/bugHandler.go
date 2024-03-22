package grpc

import (
	"bugtracker/internal/domain/bug/core/domain/aggregate"
	"bugtracker/internal/domain/bug/ports"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *grpcServer) SaveBug(ctx context.Context, in *ports.SaveBugRequest) (*ports.EmptyResponse, error) {
	s.bugService.CreateBug(ctx, aggregate.CreateBugRequest{Title: in.Title, Description: in.Description, ProjectId: in.ProjectId})
	return &ports.EmptyResponse{}, nil
}

func (s *grpcServer) GetBugById(ctx context.Context, in *ports.GetBugByIDRequest) (*ports.BugResponse, error) {
	bug, err := s.bugService.GetBugByID(ctx, in.BugId)
	if err != nil {
		return nil, err
	}

	return convertBugToBugResponse(bug), nil
}

func (s *grpcServer) GetBugsByProjectId(ctx context.Context, in *ports.GetBugsByProjectIDRequest) (*ports.GetBugsByProjectIDResponse, error) {
	bugs, err := s.bugService.GetBugsByProjectId(ctx, in.ProjectId)
	if err != nil {
		return nil, err
	}

	var convertedBugs []*ports.BugResponse
	for _, bug := range bugs {
		convertedBugs = append(convertedBugs, convertBugToBugResponse(&bug))
	}
	return &ports.GetBugsByProjectIDResponse{Bugs: convertedBugs}, nil
}

func (s *grpcServer) AssignBugTo(ctx context.Context, in *ports.AssignBugToRequest) (*ports.EmptyResponse, error) {
	err := s.bugService.AssignBugTo(ctx, in.BugId, in.Username)
	if err != nil {
		return nil, err
	}
	return &ports.EmptyResponse{}, nil
}

func convertBugToBugResponse(bug *aggregate.Bug) *ports.BugResponse {
	return &ports.BugResponse{BugId: bug.ID, Title: bug.Title, Description: bug.Description, Status: string(bug.Status), Priority: string(bug.Priority), Assignee: bug.Assignee, ProjectId: bug.ProjectId, CreatedAt: timestamppb.New(bug.CreatedAt), UpdatedAt: timestamppb.New(bug.UpdatedAt)}
}
