package registry

import (
	"go-clean-architecture/src/domain/repository"
	"go-clean-architecture/src/usecase/user"
)

type Usecase interface {
	NewGetsUser(repo repository.UserRepo) user.Gets
}

type usecaseImpl struct{}

func NewUsecase() Usecase {
	return &usecaseImpl{}
}

func (u *usecaseImpl) NewGetsUser(repo repository.UserRepo) user.Gets {
	return user.NewGets(repo)
}
