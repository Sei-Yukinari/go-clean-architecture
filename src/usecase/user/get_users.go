package user

import (
	"context"
	"go-clean-architecture/src/domain/model"
	"go-clean-architecture/src/domain/repository"
	"go-clean-architecture/src/util/appcontext"
	"go-clean-architecture/src/util/apperror"
)

type Gets interface {
	Invoke(context.Context) ([]*model.User, apperror.AppError)
}

type gets struct {
	repo repository.UserRepo
}

func NewGets(repo repository.UserRepo) Gets {
	return &gets{repo}
}

type GetsInput struct {
}

func (u *gets) Invoke(ctx context.Context) ([]*model.User, apperror.AppError) {
	users, err := u.repo.FindAllFromCache(ctx)
	logger := appcontext.GetLogger(ctx)
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	logger.Infof("users:%v\n", users)
	return users, nil
}
