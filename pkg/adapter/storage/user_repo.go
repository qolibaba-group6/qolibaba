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

var (
	ErrInvalidUserFilter = errors.New("invalid user filter")
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

func (r *userRepo) GetByFilter(ctx context.Context, filter domain.UserFilter) (*domain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	if !filter.IsValid() {
		return nil, ErrInvalidUserFilter
	}

	if  filter.ID != domain.NilUserUUID() {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.Email.IsValid() {
		q = q.Where("email = ?", filter.Email)
	}

	err := q.First(&user).Error
	
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == domain.NilUserUUID() {
		return nil, nil
	}
	
	return mapper.UserStorage2Domain(user), nil
}
