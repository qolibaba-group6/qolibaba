package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

// RedisClient is the Redis client instance
var RedisClient *redis.Client

// InitializeRedis initializes the Redis client
func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Redis server address
		DB:   0,            // Use default DB
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")
}
