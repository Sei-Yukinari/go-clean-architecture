package registry

import (
	"github.com/patrickmn/go-cache"
	"go-clean-architecture/src/domain/repository"
	"go-clean-architecture/src/gateway"
	"gorm.io/gorm"
)

type Repository interface {
	NewUser(tx *gorm.DB, cache *cache.Cache) repository.UserRepo
}

type repositoryImpl struct{}

func (r repositoryImpl) NewUser(tx *gorm.DB, cache *cache.Cache) repository.UserRepo {
	return gateway.NewUser(tx, cache)
}

func NewRepository() Repository {
	return &repositoryImpl{}
}
