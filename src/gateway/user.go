package gateway

import (
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/thoas/go-funk"
	"go-clean-architecture/src/domain/model"
	"go-clean-architecture/src/domain/repository"
	"go-clean-architecture/src/util/apperror"
	"go-clean-architecture/src/util/errorcode"
	"gorm.io/gorm"
)

type User struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewUser(db *gorm.DB, cache *cache.Cache) *User {
	return &User{db, cache}
}

var _ repository.UserRepo = (*User)(nil)

// FindAll RDBから全てのデータを取得
func (u *User) FindAll(ctx context.Context) ([]*model.User, apperror.AppError) {
	var users []*model.User

	if err := u.db.Find(&users).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			msg := fmt.Sprintf("users is not found")
			return nil, apperror.New(msg).SetCode(errorcode.NotFound).Info(msg)
		}
	}

	return users, nil
}

// FindAllFromCache インメモリキャッシュから全てのデータを取得
func (u *User) FindAllFromCache(ctx context.Context) ([]*model.User, apperror.AppError) {
	var users []*model.User
	if userCache, found := u.cache.Get(string(model.UserMaster)); found {
		users = userCache.([]*model.User)
		return users, nil
	}

	msg := fmt.Sprintf("users master is not found")
	return nil, apperror.New(msg).SetCode(errorcode.NotFoundMaster).Info(msg)
}

// Find インメモリキャッシュからfilterしてデータを取得
func (u *User) Find(ctx context.Context, id int) (*model.User, apperror.AppError) {
	users, _ := u.FindAllFromCache(ctx)
	r := funk.Find(users, func(u *model.User) bool {
		return u.ID == id
	})
	return r.(*model.User), nil
}

func (u *User) Create(ctx context.Context, user *model.User) apperror.AppError {
	err := u.db.Create(user).Error
	if err != nil {
		return apperror.Wrap(err).SetCode(errorcode.Database)
	}
	return nil
}
