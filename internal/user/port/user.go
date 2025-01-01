// internal/user/port/user.go
package port

import (
	"github.com/ehsansobhani/project_structure-3/internal/user/domain"
	"github.com/google/uuid"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	RegisterUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserProfile(id uuid.UUID) (*domain.User, error)
	UpdateUserProfile(user *domain.User) error
	DeleteUser(id uuid.UUID) error
}
