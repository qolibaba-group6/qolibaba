package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware یک middleware برای لاگ کردن درخواست‌های HTTP است
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
