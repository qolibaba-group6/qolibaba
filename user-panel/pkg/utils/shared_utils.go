
package utils

import (
	"log"
	"net/http"
	"time"
)

// MetricsMiddleware applies metrics tracking to HTTP handlers
func MetricsMiddleware(service string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)
		duration := time.Since(start).Milliseconds()

		log.Printf("[%s] %s %s %d %dms", service, r.Method, r.URL.Path, rr.statusCode, duration)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader records the status code for the response
func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

// LogError logs errors with a consistent format
func LogError(service, message string, err error) {
	log.Printf("[%s] ERROR: %s - %v", service, message, err)
}
