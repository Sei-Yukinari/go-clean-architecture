package user

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture/src/domain/model"
	"go-clean-architecture/src/gateway"
	"go-clean-architecture/src/testutil"
	"gorm.io/gorm"
	"testing"
)

func setup(t *testing.T) (context.Context, *gorm.DB, *cache.Cache, *gateway.User) {
	ctx := testutil.SetupContext()
	rdb, err := testutil.SetupRDB("user usecase", t)
	if err != nil {
		t.Errorf("error set up test rdb : %v\n", err)
	}
	inMemoryCache := testutil.SetupCache(t)
	repo := gateway.NewUser(rdb, inMemoryCache)
	return ctx, rdb, inMemoryCache, repo
}

func TestUserUsecase(t *testing.T) {
	ctx, rdb, _, userRepo := setup(t)
	t.Run("Get Users", func(t *testing.T) {
		seeds := []interface{}{
			&model.User{ID: 1, Name: "aaa"},
			&model.User{ID: 2, Name: "bbb"},
		}

		err := testutil.Seeds(rdb, seeds)
		if err != nil {
			return
		}

		users, err := NewGets(userRepo).Invoke(ctx)
		if err != nil {
			t.Errorf("usecase invoke error:%v\n", err)
		}
		assert.NoError(t, err)
		assert.Equal(t, len(users), len(seeds))
	})
}
