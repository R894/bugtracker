package repository

import "bugtracker/internal/domain/project/aggregate"

type Repository interface {
	Save(name, description, ownerId string) (*aggregate.Project, error)
	GetProjectsByUserId(userId string) ([]aggregate.Project, error)
}
