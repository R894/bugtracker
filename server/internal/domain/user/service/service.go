package service

import (
	"context"
	"database/sql"

	"github.com/R894/bugtracker/internal/domain/user/aggregate"
	"github.com/R894/bugtracker/internal/domain/user/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		userRepository: repository.NewPostgresUserRepository(db),
	}
}

func (s *UserService) CreateUser(ctx context.Context, registerRequest aggregate.UserRegisterModel) error {
	usr := aggregate.NewUser(registerRequest)
	err := s.userRepository.SaveUser(ctx, *usr)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*aggregate.User, error) {
	return s.userRepository.GetUserByID(ctx, id)
}

func (s *UserService) GetUsers(ctx context.Context) ([]aggregate.User, error) {
	return s.userRepository.GetUsers(ctx)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*aggregate.User, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, user aggregate.User) error {
	return s.userRepository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	return s.userRepository.DeleteUser(ctx, userID)
}
