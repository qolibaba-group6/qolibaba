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

type (
	UserUUID = uuid.UUID
	Email    string
)

func NilUserUUID() UserUUID {
	return UserUUID(uuid.Nil)
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

type UserListFilters struct {
	// TODO
}
