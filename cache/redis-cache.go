package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
)

//RedisCache ...
type RedisCache struct {
	host     string
	password string
	db       int
	expires  time.Duration
}

// NewRedisCache ...
func NewRedisCache(host string, password string, db int, expires time.Duration) *RedisCache {
	return &RedisCache{
		host:     host,
		password: password,
		db:       db,
		expires:  expires,
	}
}

// GetClient ...
func (cache *RedisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.password,
		DB:       cache.db,
	})
}
