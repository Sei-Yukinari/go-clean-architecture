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
	rdb, err := testutil.SetupRDB("user repository", t)
	if err != nil {
		t.Errorf("error set up test rdb : %v\n", err)
	}
	inMemoryCache := testutil.SetupCache(t)
	repo := NewUser(rdb, inMemoryCache)
	return ctx, rdb, inMemoryCache, repo
}

func TestUserRepository(t *testing.T) {
	ctx, rdb, inMemoryCache, userRepo := setup(t)
	t.Run("Get User", func(t *testing.T) {
		actual := &model.User{
			ID:   1,
			Name: "AAAAA",
		}
		seed := []interface{}{
			actual,
		}
		err := testutil.Seeds(rdb, seed)
		if err != nil {
			t.Errorf("error seed data : %v\n", err)
		}
		masterRepo := inmemorycache.MasterRepo{User: userRepo}
		inmemorycache.SetInMemoryCacheFromMaster(ctx, inMemoryCache, masterRepo)
		res, err := userRepo.Find(ctx, actual.ID)

		assert.Equal(t, err, nil)
		assert.Equal(t, res.ID, actual.ID)
		assert.Equal(t, res.Name, actual.Name)
	})
	t.Run("Create User", func(t *testing.T) {
		actual := &model.User{
			ID:   2,
			Name: "BBBB",
		}
		err := userRepo.Create(ctx, actual)
		assert.NoError(t, err)
	})
}
