package inmemorycache

import (
	"context"
	"github.com/patrickmn/go-cache"
	"go-clean-architecture/src/domain/model"
	"go-clean-architecture/src/domain/repository"
	"go-clean-architecture/src/infrastructure/logger"
)

type MasterRepo struct {
	User repository.UserRepo
}

func SetInMemoryCacheFromMaster(
	ctx context.Context,
	c *cache.Cache,
	masterRepo MasterRepo,
) error {
	// users master
	users, err := masterRepo.User.FindAll(ctx)
	if err != nil {
		logger.Fatalf("can not get user master%v\n", err)
	}
	c.Set(string(model.UserMaster), users, cache.NoExpiration)
	return nil
}
