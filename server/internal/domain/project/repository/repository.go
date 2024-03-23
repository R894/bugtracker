package repository

import (
	"context"

	"github.com/R894/bugtracker/internal/domain/project/aggregate"
)

type Repository interface {
	Save(ctx context.Context, name, description, ownerId string) (*aggregate.Project, error)
	GetById(ctx context.Context, projectId string) (*aggregate.Project, error)
	GetProjects(ctx context.Context) ([]aggregate.Project, error)
	GetProjectsByOwnerId(ctx context.Context, userId string) ([]aggregate.Project, error)
}
