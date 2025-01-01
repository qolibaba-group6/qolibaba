// pkg/adapter/storage/user_repo.go
package storage

import (
	"github.com/ehsansobhani/project_structure-3/internal/user/domain"
	"github.com/ehsansobhani/project_structure-3/internal/user/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// userRepository implements the UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}

// RegisterUser registers a new user in the database
func (r *userRepository) RegisterUser(user *domain.User) error {
	return r.db.Create(user).Error
}

// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserProfile retrieves a user profile by ID
func (r *userRepository) GetUserProfile(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserProfile updates a user's profile in the database
func (r *userRepository) UpdateUserProfile(user *domain.User) error {
	return r.db.Save(user).Error
}

// DeleteUser deletes a user from the database by ID
func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&domain.User{}, "id = ?", id).Error
}
