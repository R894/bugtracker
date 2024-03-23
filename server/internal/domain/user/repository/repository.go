package repository

import (
	"context"

	"github.com/R894/bugtracker/internal/domain/user/aggregate"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user aggregate.User) error
	GetUserByID(ctx context.Context, userID string) (*aggregate.User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*aggregate.User, error)
	GetUserByName(ctx context.Context, username string) (*aggregate.User, error)
	GetUsers(ctx context.Context) ([]aggregate.User, error)
	UpdateUser(ctx context.Context, user aggregate.User) error
	DeleteUser(ctx context.Context, userID string) error
}
