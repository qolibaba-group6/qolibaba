package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"qolibaba/pkg/conv"
	"regexp"

	"github.com/google/uuid"
)

type UserStatusType uint8

const (
	StatusUnknown UserStatusType = iota
	StatusVerified
	StatusActive
	StatusInactive
	StatusBlock
)

const (
	RoleAdmin = "admin"
	RoleUser = "user"
)

type (
	UserUUID = uuid.UUID
	Email    string
)

func NilUserUUID() UserUUID {
	return UserUUID(uuid.Nil)
}

func IsValidateUserUUID(id UserUUID) bool {
	if err := uuid.Validate(id.String()); err != nil {
		return false
	}
	return true
}

func (e Email) IsValid() bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	r := regexp.MustCompile(emailRegex)
	return r.MatchString(string(e))
}

type User struct {
	ID        UserUUID
	FirstName string
	LastName  string
	Email     Email
	Password  string
	IsAdmin   bool
	Status    UserStatusType
	Role	  string
}

func (u *User) Validate() error {
	if !u.Email.IsValid() {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func (u *User) PasswordIsCorrect(pass string) bool {
	return NewPassword(pass) == u.Password
}

func NewPassword(pass string) string {
	h := sha256.New()
	h.Write(conv.ToByte(pass))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

type UserFilter struct {
	ID UserUUID
	Email Email
}

func (f *UserFilter) IsValid() bool {
	return IsValidateUserUUID(f.ID) || f.Email.IsValid()
}
