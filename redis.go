package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	cache *redis.Client
}

func NewRedisCache(cache *redis.Client) RedisCache {
	return RedisCache{
		cache: cache,
	}
}

func (c *RedisCache) Get(key []byte) (value []byte, err error) {
	ctx := context.Background()
	val, err := c.cache.Get(ctx, string(key)).Bytes()
	return val, err
}

func (c *RedisCache) Set(key, value []byte, expireSeconds int) (err error) {
	ctx := context.Background()
	expiration := time.Duration(expireSeconds) * time.Second
	err = c.cache.Set(ctx, string(key), value, expiration).Err()
	return err
}