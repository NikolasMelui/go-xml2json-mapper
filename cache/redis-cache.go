package cache

import (
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

func newRedisCache(host string, password string, db int, expires time.Duration) *redisCache {
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

func (cache *redisCache) Set(key string, value *entity.Product) {
	client := cache.getClient()

	json, err := json.Marshal(&ProductCache{
		Data: *value,
		Hash: helper.InstanceHash(value),
	})
	if err != nil {
		panic(err)
	}
	client.Set(nil, key, json, cache.expires*time.Second)

}
