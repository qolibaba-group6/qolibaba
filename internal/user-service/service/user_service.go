
package service
import (
	"errors"
	"travel-booking-app/internal/config"
	"travel-booking-app/internal/user-service/model"
	"travel-booking-app/internal/user-service/repository"
	"travel-booking-app/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)
type UserService interface {
	Register(username, password string) (string, error)
	Login(username, password string) (string, error)
}
type userService struct {
	repo repository.UserRepository
}
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
func (s *userService) Register(username, password string) (string, error) {
	existingUser, _ := s.repo.GetUserByUsername(username)
	if existingUser != nil {
		return "", errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}
	err = s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
}
func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	token, err := utils.GenerateToken(user.ID.String(), config.AppConfig.JWT.Secret, config.AppConfig.JWT.Expiry)
	if err != nil {
		return "", err
	}
	return token, nil
}
