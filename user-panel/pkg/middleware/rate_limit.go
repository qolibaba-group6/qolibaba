
package middleware

import (
	"net/http"
	"time"
	"user-panel/pkg/cache"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		// Get the current count from Redis
		count, err := cache.RedisClient.Get(cache.ctx, key).Int()
		if err != nil && err.Error() != "redis: nil" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit check failed"})
			c.Abort()
			return
		}

		// If the count exceeds the limit, block the request
		if count >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		// Increment the count and set expiration if it's the first request
		err = cache.RedisClient.Incr(cache.ctx, key).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit increment failed"})
			c.Abort()
			return
		}
		cache.RedisClient.Expire(cache.ctx, key, duration)

		// Allow the request
		c.Next()
	}
}
