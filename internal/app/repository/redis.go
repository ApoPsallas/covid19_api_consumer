package repository

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type CacheMapper struct {
	Client redis.Client
}

func (r CacheMapper) Exists(key string) (int64, error) {
	return r.Client.Exists(key).Result()
}

func (r CacheMapper) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

func (r CacheMapper) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.Client.Set(key, value, expiration).Result()
}
