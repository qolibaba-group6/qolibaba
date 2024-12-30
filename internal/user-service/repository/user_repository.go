
package repository
import (
	"travel-booking-app/internal/user-service/model"
	"travel-booking-app/internal/database"
)
type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
}
type userRepository struct{}
func NewUserRepository() UserRepository {
	return &userRepository{}
}
func (r *userRepository) CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}
func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
