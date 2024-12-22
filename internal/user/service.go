package user

import (
	"context"
	"errors"
	"qolibaba/internal/user/domain"
	"qolibaba/internal/user/port"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("validation failed")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidEmail           = errors.New("invalid email format")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, user domain.User) (domain.UserUUID, error) {
	if ok := user.Email.IsValid(); !ok {
		return domain.NilUserUUID(), ErrInvalidEmail
	}

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		// log
		return domain.NilUserUUID(), err
	}

	return userID, nil
}

func (s *service) GetByID(ctx context.Context, userID domain.UserUUID) (*domain.User, error) {
	panic("not implemented")
}
