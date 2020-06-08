package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nikolasmelui/go-xml2json-mapper/entity"
	"github.com/nikolasmelui/go-xml2json-mapper/helper"
)

type redisCache struct {
	host     string
	password string
	db       int
	expires  time.Duration
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

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.password,
		DB:       cache.db,
	})
}

// Set ...
func (cache *redisCache) Set(key string, value *entity.Product) {
	client := cache.getClient()

	json, err := json.Marshal(&ProductWithHash{
		Data: *value,
		Hash: helper.InstanceHash(value),
	})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Set(ctx, key, json, cache.expires*time.Second)

}

func (cache *redisCache) Get(key string) *ProductWithHash {

	client := cache.getClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	productWithHash := ProductWithHash{}
	err = json.Unmarshal([]byte(val), &productWithHash)
	if err != nil {
		panic(err)
	}

	return &productWithHash
}
