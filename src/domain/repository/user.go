package repository

import (
	"context"
	"go-clean-architecture/src/domain/model"
	"go-clean-architecture/src/util/apperror"
)

type UserRepo interface {
	FindAll(ctx context.Context) ([]*model.User, apperror.AppError)
	FindAllFromCache(ctx context.Context) ([]*model.User, apperror.AppError)
	Find(ctx context.Context, id int) (*model.User, apperror.AppError)
	Create(context.Context, *model.User) apperror.AppError
}
