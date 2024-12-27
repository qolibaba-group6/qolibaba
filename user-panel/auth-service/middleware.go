
package main

import (
    "net/http"
    "strings"
)

// AuthMiddleware checks the validity of the JWT token in the request header
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        // Extract the token from the header
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
            return
        }

        // Validate the token
        userID, err := ValidateToken(tokenString)
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }

        // Add userID to the request context
        ctx := r.Context()
        ctx = context.WithValue(ctx, "userID", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
