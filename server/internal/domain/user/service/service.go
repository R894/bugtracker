package service

import (
	"bugtracker/internal/domain/user/aggregate"
	"bugtracker/internal/domain/user/repository"
	"context"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username, email, password string) error {
	usr := aggregate.NewUser(username, email, password)
	err := s.userRepository.SaveUser(ctx, *usr)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*aggregate.User, error) {
	return s.userRepository.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user aggregate.User) error {
	return s.userRepository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	return s.userRepository.DeleteUser(ctx, userID)
}
