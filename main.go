package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/infrastructure/cache"
	"go-clean-architecture/src/infrastructure/logger"
	"go-clean-architecture/src/infrastructure/rdb"
	"go-clean-architecture/src/infrastructure/redis"
	"go-clean-architecture/src/infrastructure/server"
	"go-clean-architecture/src/interfaces/controller"
	"go-clean-architecture/src/interfaces/controller/validation"
	"go-clean-architecture/src/middleware"
	"go-clean-architecture/src/registry"
	"go-clean-architecture/src/util/inmemorycache"
)

func main() {

	rdb := rdb.NewRDB()
	redis := redis.NewRedis()
	middlewares := []gin.HandlerFunc{
		middleware.NewCors(),
		middleware.NewInjectLogger(),
		middleware.NewInjectProfile(),
	}

	repo := registry.NewRepository()
	ctx := context.Background()

	inmemoryCache := cache.NewInmemoryCache()

	masterRepo := inmemorycache.MasterRepo{
		User: repo.NewUser(rdb, inmemoryCache),
	}

	err := inmemorycache.SetInMemoryCacheFromMaster(
		ctx,
		inmemoryCache,
		masterRepo,
	)
	if err != nil {
		logger.Fatalf("failed set inmemory cache from master data")
	}

	ctrl := controller.NewController(
		rdb,
		inmemoryCache,
		redis,
		repo,
		registry.NewUsecase(),
		validation.NewValidator(),
	)

	router := server.NewRouter(middlewares, ctrl)
	server.Run(router)
}
