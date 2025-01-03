// pkg/cache/redis.go
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache implements the Cache interface using Redis
type RedisCache struct {
	Client *redis.Client
}

// NewRedisCache creates a new RedisCache instance
func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisCache{Client: rdb}
}

// Set sets a value in Redis with a specified expiration duration in seconds
func (c *RedisCache) Set(key string, value interface{}, expiration int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.Client.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
}

// Get retrieves a value from Redis by key
func (c *RedisCache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.Client.Get(ctx, key).Result()
}

// Delete removes a value from Redis by key
func (c *RedisCache) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.Client.Del(ctx, key).Err()
}
