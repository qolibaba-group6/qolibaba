
package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

// TestLoginEndpoint tests the /login endpoint
func TestLoginEndpoint(t *testing.T) {
    reqBody := `{"user_id": "user1", "password": "password123"}`
    req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(LoginHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Fatalf("Expected status code %v, got %v", http.StatusOK, status)
    }

    responseBody := rr.Body.String()
    if !strings.Contains(responseBody, "token") {
        t.Fatalf("Expected response to contain token, got %v", responseBody)
    }
}

// TestProtectedEndpoint tests the /protected endpoint with a valid token
func TestProtectedEndpoint(t *testing.T) {
    // Generate a valid token
    token, err := GenerateToken("user1")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    req := httptest.NewRequest(http.MethodGet, "/protected", nil)
    req.Header.Set("Authorization", "Bearer "+token)

    rr := httptest.NewRecorder()
    handler := AuthMiddleware(http.HandlerFunc(ProtectedHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Fatalf("Expected status code %v, got %v", http.StatusOK, status)
    }

    responseBody := rr.Body.String()
    if !strings.Contains(responseBody, "user_id") {
        t.Fatalf("Expected response to contain user_id, got %v", responseBody)
    }
}

// TestProtectedEndpointWithoutToken tests the /protected endpoint without a token
func TestProtectedEndpointWithoutToken(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/protected", nil)
    rr := httptest.NewRecorder()
    handler := AuthMiddleware(http.HandlerFunc(ProtectedHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusUnauthorized {
        t.Fatalf("Expected status code %v, got %v", http.StatusUnauthorized, status)
    }
}
