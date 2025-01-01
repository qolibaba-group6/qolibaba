// pkg/middleware/logging.go
package middleware

import (
	"net/http"
	"time"

	"github.com/ehsansobhani/travel_agencies/pkg/logger"
)

// LoggingMiddleware logs each incoming HTTP request
func LoggingMiddleware(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// پاسخ‌دهی با استفاده از ResponseWriter سفارشی برای گرفتن وضعیت پاسخ
			lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(lrw, r)

			duration := time.Since(start)
			log.Infof("Method: %s, URI: %s, Status: %d, Duration: %v",
				r.Method, r.RequestURI, lrw.statusCode, duration)
		})
	}
}

// loggingResponseWriter wraps http.ResponseWriter to capture the status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
