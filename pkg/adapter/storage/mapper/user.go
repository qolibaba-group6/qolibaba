package mapper

import (
	"qolibaba/internal/user/domain"
	"qolibaba/pkg/adapter/storage/types"
)

func UserDomain2Storage(user domain.User) *types.User {
	return &types.User{
		Model: types.Model{
			ID: user.ID,
		},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     string(user.Email),
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		Status:    uint8(user.Status),
		Role:      user.Role,
	}
}

func UserStorage2Domain(user types.User) *domain.User {
	return &domain.User{
		ID:        domain.UserUUID(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     domain.Email(user.Email),
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		Status:    domain.UserStatusType(user.Status),
		Role:      user.Role,
	}
}
