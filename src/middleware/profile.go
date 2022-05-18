package middleware

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/util/appcontext"
	"time"
)

func NewInjectProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		logger := appcontext.GetLogger(c.Request.Context())
		c.Next()

		latency := time.Since(t)
		logger.Infof("response time(%s): %+vms elapsed", c.Request.URL, latency)
	}
}
