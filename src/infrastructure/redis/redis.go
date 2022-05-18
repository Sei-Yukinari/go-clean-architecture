package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-clean-architecture/src/config"
	"go-clean-architecture/src/infrastructure/logger"
	"runtime"
	"time"
)

type Client = redis.Client

var Nil = redis.Nil

func NewRedis() *Client {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Conf.Redis.URL,
		Password:     config.Conf.Redis.Password,
		DB:           config.Conf.Redis.DB,
		PoolSize:     runtime.NumCPU() * config.Conf.Redis.PoolSizePerCPU,
		MinIdleConns: config.Conf.Redis.MinIdleConnection,
		PoolTimeout:  time.Duration(config.Conf.Redis.PoolTimeoutSeconds) * time.Second,
	})

	ctx := context.Background()
	err := client.Ping(ctx).Err()

	if err != nil {
		logger.Fatalf("failed to connect redis:%v\n", err)
	}

	logger.Info("success to connect redis!")

	return client
}
