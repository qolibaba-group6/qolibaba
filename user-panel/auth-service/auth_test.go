
package main

import (
    "testing"
    "time"
)

// TestGenerateToken tests the GenerateToken function
func TestGenerateToken(t *testing.T) {
    userID := "testuser"
    token, err := GenerateToken(userID)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if token == "" {
        t.Fatalf("Expected a token, got an empty string")
    }
}

// TestValidateToken tests the ValidateToken function
func TestValidateToken(t *testing.T) {
    userID := "testuser"
    token, err := GenerateToken(userID)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    validatedUserID, err := ValidateToken(token)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if validatedUserID != userID {
        t.Fatalf("Expected userID %v, got %v", userID, validatedUserID)
    }
}

// TestExpiredToken tests the behavior with an expired token
func TestExpiredToken(t *testing.T) {
    expiredClaims := &Claims{
        UserID: "testuser",
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    _, err = ValidateToken(tokenString)
    if err == nil {
        t.Fatalf("Expected an error for expired token, got none")
    }
}
