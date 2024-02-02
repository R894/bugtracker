package service

import (
	"bugtracker/internal/domain/project/aggregate"
	"bugtracker/internal/domain/project/repository"
)

type ProjectService struct {
	repo repository.Repository
}

func NewProjectService(projectRepo repository.Repository) *ProjectService {
	return &ProjectService{
		repo: projectRepo,
	}
}

func (s *ProjectService) CreateProject(name, description, ownerId string) (*aggregate.Project, error) {
	project, err := s.repo.Save(name, description, ownerId)
	if err != nil {
		return nil, err
	}
	return project, nil
}
