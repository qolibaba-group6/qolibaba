// pkg/middleware/rate_limit.go
package middleware

import (
	"net/http"
	"time"

	"github.com/ehsansobhani/travel_agencies/pkg/logger"
	"golang.org/x/time/rate"
)

// RateLimitMiddleware limits the number of requests per duration per IP
func RateLimitMiddleware(limit int, duration time.Duration, log *logger.Logger) func(http.Handler) http.Handler {
	// ایجاد یک نرخ محدودکننده برای هر IP
	visitors := make(map[string]*rate.Limiter)

	// کانال برای پاکسازی مداوم IPهای قدیمی
	go func() {
		for {
			time.Sleep(time.Minute)
			for ip, limiter := range visitors {
				if limiter.AllowN(time.Now(), 1) {
					delete(visitors, ip)
					log.Infof("Removed IP %s from rate limiter", ip)
				}
			}
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getIP(r)
			limiter, exists := visitors[ip]
			if !exists {
				limiter = rate.NewLimiter(rate.Every(duration/time.Duration(limit)), limit)
				visitors[ip] = limiter
			}

			if !limiter.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// getIP extracts the client's real IP address from the request
func getIP(r *http.Request) string {
	// ابتدا بررسی کنید که آیا درخواست از طریق پروکسی ارسال شده است
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
