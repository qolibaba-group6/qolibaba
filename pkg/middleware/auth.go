// pkg/middleware/auth.go
package middleware

import (
	"context"
	"net/http"

	"github.com/ehsansobhani/project_structure-3/pkg/utils"
)

// Define a custom type for context keys to avoid collisions
type contextKey string

const userIDKey contextKey = "userID"

// AuthMiddleware validates JWT and adds user ID to context
func AuthMiddleware(next http.HandlerFunc, jwtSecret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(tokenStr, jwtSecret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper function to retrieve user ID from context
func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}
