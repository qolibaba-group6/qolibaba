package port

import (
	"context"
	"qolibaba/internal/user/domain"
)

type Service interface {
	Create(ctx context.Context, user domain.User) (domain.UserUUID, error)
	GetByID(ctx context.Context, userID domain.UserUUID) (*domain.User, error)
	GetByFilter(ctx context.Context, filter domain.UserFilter) (*domain.User, error)
}
