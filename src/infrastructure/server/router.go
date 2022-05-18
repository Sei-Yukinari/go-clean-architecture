package server

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/config"
	"go-clean-architecture/src/interfaces/controller"
	"net/http"
)

func NewRouter(
	middlewares []gin.HandlerFunc,
	ctrl *controller.Controller,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	if config.IsTest() {
		gin.SetMode(gin.TestMode)
	}
	if config.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	for _, m := range middlewares {
		r.Use(m)
	}

	healthRouter(r)

	apiV1 := r.Group("/api/v1")

	userRouter(apiV1, ctrl)

	return r
}

// health check API
func healthRouter(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}

func userRouter(rg *gin.RouterGroup, ctrl *controller.Controller) {
	controller := controller.NewUserController(ctrl)
	rg.GET("/", func(c *gin.Context) {
		controller.Find(c)
	})
	rg.POST("/", func(c *gin.Context) {
		controller.Find(c)
	})
}
