package repository

import (
	"context"
	"errors"

	"github.com/R894/bugtracker/internal/domain/bug/core/domain/aggregate"
)

type InMemoryBugRepository struct {
	bugs map[string]*aggregate.Bug
}

func NewInMemoryBugRepository() *InMemoryBugRepository {
	return &InMemoryBugRepository{
		bugs: make(map[string]*aggregate.Bug),
	}
}

func (r *InMemoryBugRepository) SaveBug(ctx context.Context, bug *aggregate.Bug) error {
	r.bugs[bug.ID] = bug
	return nil
}

func (r *InMemoryBugRepository) GetBugByID(ctx context.Context, bugID string) (*aggregate.Bug, error) {
	bug, ok := r.bugs[bugID]
	if !ok {
		return nil, errors.New("bug not found")
	}
	return bug, nil
}

func (r *InMemoryBugRepository) GetBugsByProjectID(ctx context.Context, projectID string) ([]aggregate.Bug, error) {
	var bugs []aggregate.Bug
	for _, bug := range r.bugs {
		if bug.ProjectId == projectID {
			bugs = append(bugs, *bug)
		}
	}
	return bugs, nil
}

func (r *InMemoryBugRepository) GetBugs(ctx context.Context) ([]aggregate.Bug, error) {
	var bugs []aggregate.Bug
	for _, bug := range r.bugs {
		bugs = append(bugs, *bug)
	}
	return bugs, nil
}

func (r *InMemoryBugRepository) AssignBugTo(ctx context.Context, bugID, username string) error {
	bug, ok := r.bugs[bugID]
	if !ok {
		return errors.New("bug not found")
	}
	bug.Assignee = username
	r.bugs[bugID] = bug
	return nil
}

func (r *InMemoryBugRepository) UpdateBug(ctx context.Context, update *aggregate.UpdateBugRequest) error {
	bug, ok := r.bugs[update.ID]
	if !ok {
		return errors.New("bug not found")
	}
	bug.Status = update.Status
	bug.Description = update.Description
	r.bugs[update.ID] = bug
	return nil
}

func (r *InMemoryBugRepository) DeleteBug(ctx context.Context, bugID string) error {
	if _, ok := r.bugs[bugID]; !ok {
		return errors.New("bug not found")
	}
	delete(r.bugs, bugID)
	return nil
}
