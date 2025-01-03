// pkg/cache/cache.go
package cache

// Cache defines the interface for caching operations
type Cache interface {
	// Set sets a value in the cache with a specified expiration duration in seconds
	Set(key string, value interface{}, expiration int) error

	// Get retrieves a value from the cache by key
	Get(key string) (string, error)

	// Delete removes a value from the cache by key
	Delete(key string) error
}
