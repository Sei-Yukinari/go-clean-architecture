package testutil

import (
	gocache "github.com/patrickmn/go-cache"
	"go-clean-architecture/src/infrastructure/cache"
	"testing"
)

func SetupCache(t *testing.T) *gocache.Cache {
	t.Helper()
	inmemoryCache := cache.NewInmemoryCache()
	return inmemoryCache
}
