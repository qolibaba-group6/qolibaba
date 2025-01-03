package repository

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// User مخزن داده‌های کاربر
type User struct {
	ID       uuid.UUID
	Email    string
	Password string
}

// UserRepository اینترفیس برای مخزن کاربران
type UserRepository interface {
	FindUserByEmail(email string) (User, error)
	SaveUser(user User) error
}

// مخزن کاربران در حافظه
type userRepository struct {
	users map[string]User
}

// NewUserRepository ایجاد یک نمونه از UserRepository
func NewUserRepository() UserRepository {
	return &userRepository{
		users: make(map[string]User),
	}
}

// FindUserByEmail جستجوی کاربر بر اساس ایمیل
func (r *userRepository) FindUserByEmail(email string) (User, error) {
	user, exists := r.users[email]
	if !exists {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// SaveUser ذخیره کاربر در حافظه
func (r *userRepository) SaveUser(user User) error {
	r.users[user.Email] = user
	return nil
}
