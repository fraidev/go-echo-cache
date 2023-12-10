package cache

import "github.com/coocood/freecache"

type MemoryCache struct {
	cache *freecache.Cache
}

func NewMemoryCache(cache *freecache.Cache) MemoryCache {
	return MemoryCache{
		cache: cache,
	}
}

func (c *MemoryCache) Get(key []byte) (value []byte, err error) {
	return c.cache.Get(key)
}

func (c *MemoryCache) Set(key, value []byte, expireSeconds int) (err error) {
	return c.cache.Set(key, value, expireSeconds)
}
