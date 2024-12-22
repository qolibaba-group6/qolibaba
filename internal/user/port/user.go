package port

import (
	"context"
	"qolibaba/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserUUID, error)
	GetByID(ctx context.Context, userID domain.UserUUID) (*domain.User, error)
	Get(ctx context.Context, filter domain.UserListFilters) ([]domain.User, error)
}
