package service

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/user/aggregate"
	"bugtracker/internal/domain/user/repository"
	"context"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(db *database.PostgresDB) *UserService {
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

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*aggregate.User, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, user aggregate.User) error {
	return s.userRepository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	return s.userRepository.DeleteUser(ctx, userID)
}
