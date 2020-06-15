package cache

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host     string
	password string
	db       int
	expires  time.Duration
}

// RedisClient ...
type RedisClient struct {
	*redis.Client
}

// NewRedisCache ...
func NewRedisCache(host string, password string, db int, expires time.Duration) ProductCache {
	return &redisCache{
		host:     host,
		password: password,
		db:       db,
		expires:  expires,
	}
}

var once sync.Once
var redisClient *RedisClient

func (cache *redisCache) getClient() *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     cache.host,
			Password: cache.password,
			DB:       cache.db,
			PoolSize: 10,
		})
		redisClient = &RedisClient{client}
	})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to the redis %v", err)
	}
	return redisClient

}

// Set ...
func (cache *redisCache) Set(key string, value *ProductWithHash) {
	client := cache.getClient()

	json, err := json.Marshal(&value)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client.Set(ctx, key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *ProductWithHash {

	client := cache.getClient()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	val, err := client.Get(ctx, key).Result()
	// TODO: Custom error and handle need (return the error)
	if err != nil {
		return nil
	}

	productWithHash := &ProductWithHash{}
	err = json.Unmarshal([]byte(val), &productWithHash)

	return productWithHash
}
