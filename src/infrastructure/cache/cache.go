package cache

import "github.com/patrickmn/go-cache"

type InMemoryCache struct {
	cache *cache.Cache
}

var c *cache.Cache

func NewInmemoryCache() *cache.Cache {
	if c != nil {
		return c
	}
	newCache := cache.New(cache.NoExpiration, cache.NoExpiration)
	c = newCache
	return c
}
