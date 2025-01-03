// internal/user/port/service.go
package port

import (
	"github.com/ehsansobhani/project_structure-3/internal/user/domain"
	"github.com/google/uuid"
)


type Service interface {
	Create(ctx context.Context, user domain.User) (domain.UserUUID, error)
	GetByID(ctx context.Context, userID domain.UserUUID) (*domain.User, error)
	GetByFilter(ctx context.Context, filter domain.UserFilter) (*domain.User, error)
	UpdateRole(ctx context.Context, userID domain.UserUUID, role string) error
}
