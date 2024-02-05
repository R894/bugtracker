package service

import (
	"bugtracker/internal/domain/project/aggregate"
	"bugtracker/internal/domain/project/repository"
	"context"
	"database/sql"
)

type ProjectService struct {
	repo repository.Repository
}

func NewProjectService(db *sql.DB) *ProjectService {
	return &ProjectService{
		repo: repository.NewPostgresRepository(db),
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, name, description, ownerId string) (*aggregate.Project, error) {
	project, err := s.repo.Save(ctx, name, description, ownerId)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *ProjectService) GetProjectByID(ctx context.Context, projectId string) (*aggregate.Project, error) {
	return s.repo.GetById(ctx, projectId)
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]aggregate.Project, error) {
	return s.repo.GetProjects(ctx)
}
