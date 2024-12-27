
package main

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "errors"
)

// Secret key for signing JWTs
var jwtSecret = []byte("your_secret_key")

// Claims struct for JWT payload
type Claims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

// GenerateToken generates a new JWT for a given user ID
func GenerateToken(userID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateToken validates a JWT and extracts the user ID
func ValidateToken(tokenString string) (string, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
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

    return claims.UserID, nil
}
