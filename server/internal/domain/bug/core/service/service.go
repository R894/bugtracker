package service

import (
	"context"
	"database/sql"

	"github.com/R894/bugtracker/internal/domain/bug/adapters/repository"
	"github.com/R894/bugtracker/internal/domain/bug/core/domain/aggregate"
	"github.com/R894/bugtracker/internal/domain/bug/ports"
)

type BugService struct {
	bugRepository ports.BugRepository
}

func NewBugService(db *sql.DB) *BugService {
	return &BugService{
		bugRepository: repository.NewPostgresBugRepository(db),
	}
}

func (s *BugService) CreateBug(ctx context.Context, req aggregate.CreateBugRequest) (*aggregate.Bug, error) {
	bug := aggregate.NewBug(req)
	err := s.bugRepository.SaveBug(ctx, bug)
	if err != nil {
		return nil, err
	}
	return bug, nil
}

func (s *BugService) GetBugs(ctx context.Context) ([]aggregate.Bug, error) {
	bugs, err := s.bugRepository.GetBugs(ctx)
	if err != nil {
		return nil, err
	}

	return bugs, nil
}

func (s *BugService) GetBugByID(ctx context.Context, bugID string) (*aggregate.Bug, error) {
	return s.bugRepository.GetBugByID(ctx, bugID)
}

func (s *BugService) GetBugsByProjectId(ctx context.Context, projectId string) ([]aggregate.Bug, error) {
	return s.bugRepository.GetBugsByProjectID(ctx, projectId)
}

func (s *BugService) UpdateBug(ctx context.Context, bug *aggregate.UpdateBugRequest) error {
	return s.bugRepository.UpdateBug(ctx, bug)
}

func (s *BugService) DeleteBug(ctx context.Context, bugID string) error {
	return s.bugRepository.DeleteBug(ctx, bugID)
}

func (s *BugService) AssignBugTo(ctx context.Context, bugId, username string) error {
	return s.bugRepository.AssignBugTo(ctx, bugId, username)
}
