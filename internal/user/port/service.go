// internal/user/port/service.go
package port

import (
	"github.com/ehsansobhani/project_structure-3/internal/user/domain"
	"github.com/google/uuid"
)

// UserService defines the interface for user service operations
type UserService interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
	GetProfile(id uuid.UUID) (*domain.User, error)
	UpdateProfile(user *domain.User) error
	DeleteUser(id uuid.UUID) error
}
