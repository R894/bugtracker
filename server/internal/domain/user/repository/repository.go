package repository

import (
	"bugtracker/internal/domain/user/aggregate"
	"context"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user aggregate.User) error
	GetUserByID(ctx context.Context, userID string) (*aggregate.User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*aggregate.User, error)
	GetUsers(ctx context.Context) ([]aggregate.User, error)
	UpdateUser(ctx context.Context, user aggregate.User) error
	DeleteUser(ctx context.Context, userID string) error
}
