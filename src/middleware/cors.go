package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/config"
	"strings"
	"time"
)

func NewCors() gin.HandlerFunc {
	AllowOrigins := strings.Split(config.Conf.App.CorsOrigins, ",")
	return cors.New(cors.Config{
		AllowOrigins:     AllowOrigins,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
