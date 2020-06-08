package cache

import (
	"time"
)

//RedisCache ...
type RedisCache struct {
	host     string
	port     int
	db       string
	password string
	expires  time.Duration
}

// NewRedisCache ...
func NewRedisCache(host string, port int, db string, password string, expires time.Duration) *RedisCache {
	return &RedisCache{
		host:     host,
		port:     port,
		db:       db,
		password: password,
		expires:  expires,
	}
}
