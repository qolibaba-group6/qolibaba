package storage

import (
	"context"
	"errors"
	"qolibaba/internal/user/domain"
	userPort "qolibaba/internal/user/port"
	"qolibaba/pkg/adapter/storage/mapper"
	"qolibaba/pkg/adapter/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPort.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user domain.User) (domain.UserUUID, error) {
	u := mapper.UserDomain2Storage(user)
	return domain.UserUUID(u.ID), r.db.WithContext(ctx).Table("users").Create(u).Error
}

func (r *userRepo) GetByID(ctx context.Context, userID domain.UserUUID) (*domain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	err := q.Where("id = ?", userID).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}

func (r *userRepo) Get(ctx context.Context, filter domain.UserListFilters) ([]domain.User, error) {
	panic("not implemented")
}
