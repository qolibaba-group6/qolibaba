
package main

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "errors"
)

// Secret keys for signing JWTs
var jwtAccessSecret = []byte("your_access_secret")
var jwtRefreshSecret = []byte("your_refresh_secret")

// Extended Claims struct for Role-Based Access Control (RBAC)
type ExtendedClaims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"` // e.g., "admin", "user"
    jwt.StandardClaims
}

// GenerateAccessToken generates a new access token with role information
func GenerateAccessToken(userID, role string) (string, error) {
    expirationTime := time.Now().Add(15 * time.Minute) // Access token expires in 15 minutes
    claims := &ExtendedClaims{
        UserID: userID,
        Role:   role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtAccessSecret)
}

// GenerateRefreshToken generates a new refresh token
func GenerateRefreshToken(userID string) (string, error) {
    expirationTime := time.Now().Add(7 * 24 * time.Hour) // Refresh token expires in 7 days
    claims := &jwt.StandardClaims{
        Subject:   userID,
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtRefreshSecret)
}

// ValidateAccessToken validates an access token and extracts the claims
func ValidateAccessToken(tokenString string) (*ExtendedClaims, error) {
    claims := &ExtendedClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtAccessSecret, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, errors.New("invalid token signature")
        }
        return nil, errors.New("could not parse token")
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}

// ValidateRefreshToken validates a refresh token and extracts the user ID
func ValidateRefreshToken(tokenString string) (string, error) {
    claims := &jwt.StandardClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtRefreshSecret, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return "", errors.New("invalid token signature")
        }
        return "", errors.New("could not parse token")
    }

    if !token.Valid {
        return "", errors.New("invalid token")
    }

    return claims.Subject, nil
}
