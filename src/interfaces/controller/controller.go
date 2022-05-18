package controller

import (
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"go-clean-architecture/src/interfaces/controller/validation"
	"go-clean-architecture/src/interfaces/presenter"
	"go-clean-architecture/src/registry"
	"gorm.io/gorm"
)

type Controller struct {
	db            *gorm.DB
	inMemoryCache *cache.Cache
	redisClient   *redis.Client
	repo          registry.Repository
	usecase       registry.Usecase
	presenter     *presenter.Presenter
	validator     validation.Validator
}

func NewController(
	db *gorm.DB,
	inMemoryCache *cache.Cache,
	redisClient *redis.Client,
	repo registry.Repository,
	usecase registry.Usecase,
	validator validation.Validator,
) *Controller {
	return &Controller{
		db:            db,
		inMemoryCache: inMemoryCache,
		redisClient:   redisClient,
		repo:          repo,
		usecase:       usecase,
		presenter:     &presenter.Presenter{},
		validator:     validator,
	}
}
