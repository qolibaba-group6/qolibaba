package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"travel-booking-app/internal/config"
	"travel-booking-app/pkg/utils"

	"github.com/golang-jwt/jwt"
)

// AuthMiddleware بررسی می‌کند که آیا توکن معتبر است یا خیر
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		// پارس توکن با استفاده از ParseWithClaims
		token, err := jwt.ParseWithClaims(tokenStr, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWT.Secret), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// بررسی معتبر بودن توکن
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// استخراج claims
		claims, ok := token.Claims.(*utils.Claims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// اگر نیاز دارید UserID را چاپ کنید یا در context قرار دهید، از آن استفاده کنید:
		fmt.Println("UserID from token:", claims.UserID)

		// ادامه زنجیره
		next.ServeHTTP(w, r)
	})
}
