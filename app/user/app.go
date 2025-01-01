// app/user/app.go
package user

import (
	"github.com/ehsansobhani/project_structure-3/internal/user/domain"
	"github.com/ehsansobhani/project_structure-3/internal/user/port"
	"github.com/ehsansobhani/project_structure-3/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserService implements the UserService interface
type UserService struct {
	repo      port.UserRepository
	jwtSecret []byte
}

// NewUserService creates a new UserService
func NewUserService(repo port.UserRepository, jwtSecret []byte) port.UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

// Register registers a new user
func (s *UserService) Register(user *domain.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.ID = uuid.New()
	return s.repo.RegisterUser(user)
}

// Login logs in a user and returns a JWT token
func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.String(), s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetProfile retrieves a user's profile
func (s *UserService) GetProfile(id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserProfile(id)
}

// UpdateProfile updates a user's profile
func (s *UserService) UpdateProfile(user *domain.User) error {
	if user.Password != "" {
		// Hash the new password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return s.repo.UpdateUserProfile(user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
