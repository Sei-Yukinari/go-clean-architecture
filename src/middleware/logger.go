package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-clean-architecture/src/infrastructure/logger"
	"go-clean-architecture/src/util/appcontext"
)

// https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-request-tracing.html
// X-Amzn-Trace-Id: Root=1-67891233-abcdef012345678912345678
const AlbRequestIdHeader string = "x-amzn-trace-id"

func NewInjectLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(AlbRequestIdHeader)
		logger := logger.New().WithFields(logrus.Fields{"requestId": requestID})

		newCtx := appcontext.SetLogger(c.Request.Context(), logger)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
