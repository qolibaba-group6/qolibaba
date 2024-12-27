
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestWalletBalanceEndpoint tests the /wallet/balance endpoint
func TestWalletBalanceEndpoint(t *testing.T) {
    // Generate a valid token
    token, err := GenerateToken("user1")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    req := httptest.NewRequest(http.MethodGet, "/wallet/balance", nil)
    req.Header.Set("Authorization", "Bearer "+token)

    rr := httptest.NewRecorder()
    handler := AuthMiddleware(http.HandlerFunc(WalletBalanceHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Fatalf("Expected status code %v, got %v", http.StatusOK, status)
    }

    responseBody := rr.Body.String()
    if !strings.Contains(responseBody, "user_id") {
        t.Fatalf("Expected response to contain user_id, got %v", responseBody)
    }
}
