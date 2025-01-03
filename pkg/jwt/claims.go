package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserID  uuid.UUID
	// IsAdmin bool
	Role    string
}
