package gateway

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture/src/domain/model"
	"go-clean-architecture/src/testutil"
	"go-clean-architecture/src/util/inmemorycache"
	"gorm.io/gorm"
	"testing"
)

func setup(t *testing.T) (context.Context, *gorm.DB, *cache.Cache, *User) {
	ctx := context.Background()
	rdb := testutil.SetupRDB(t)
	inMemoryCache := testutil.SetupCache(t)
	repo := NewUser(rdb, inMemoryCache)
	return ctx, rdb, inMemoryCache, repo
}

func TestUserRepository(t *testing.T) {
	ctx, _, inMemoryCache, userRepo := setup(t)
	t.Run("Create and Get", func(t *testing.T) {
		actual := &model.User{
			ID:   1,
			Name: "AAAAA",
		}
		userRepo.Create(ctx, actual)
		masterRepo := inmemorycache.MasterRepo{User: userRepo}
		inmemorycache.SetInMemoryCacheFromMaster(ctx, inMemoryCache, masterRepo)
		res, err := userRepo.Find(ctx, actual.ID)

		assert.Equal(t, err, nil)
		assert.Equal(t, res.ID, actual.ID)
		assert.Equal(t, res.Name, actual.Name)
	})
}
