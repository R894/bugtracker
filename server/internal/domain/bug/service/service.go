package service

import (
	"bugtracker/internal/domain/bug/aggregate"
	"bugtracker/internal/domain/bug/repository"
	"context"
)

type BugService struct {
	bugRepository repository.BugRepository
}

func NewBugService(bugRepository repository.BugRepository) *BugService {
	return &BugService{
		bugRepository: bugRepository,
	}
}

func (s *BugService) CreateBug(ctx context.Context, title, description string) (*aggregate.Bug, error) {
	bug := aggregate.NewBug(title, description)
	err := s.bugRepository.SaveBug(ctx, bug)
	if err != nil {
		return nil, err
	}
	return bug, nil
}

func (s *BugService) GetBugByID(ctx context.Context, bugID string) (*aggregate.Bug, error) {
	return s.bugRepository.GetBugByID(ctx, bugID)
}

func (s *BugService) UpdateBug(ctx context.Context, bug *aggregate.Bug) error {
	return s.bugRepository.UpdateBug(ctx, bug)
}

func (s *BugService) DeleteBug(ctx context.Context, bugID string) error {
	return s.bugRepository.DeleteBug(ctx, bugID)
}
