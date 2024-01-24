package domain

import "context"

type UserRepository interface {
	SaveUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, userID string) (*User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID string) error
}
