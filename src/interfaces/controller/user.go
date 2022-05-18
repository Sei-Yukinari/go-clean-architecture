package controller

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/infrastructure/http"
	"go-clean-architecture/src/util/apperror"
	netHttp "net/http"
)

type UserController struct {
	controller *Controller
}

func NewUserController(controller *Controller) *UserController {
	return &UserController{
		controller: controller,
	}
}

func (u *UserController) Find(c *gin.Context) {
	ctx := c.Request.Context()
	userRepo := u.controller.repo.NewUser(u.controller.db, u.controller.inMemoryCache)
	users, err := u.
		controller.
		usecase.
		NewGetsUser(userRepo).
		Invoke(ctx)
	if err != nil {
		http.NewError(
			c,
			netHttp.StatusInternalServerError,
			apperror.Wrap(err).Info("failed to get users"))
		return
	}
	http.NewResponse(
		c,
		netHttp.StatusOK,
		"aaa",
		u.controller.presenter.Users(users),
	)
}
