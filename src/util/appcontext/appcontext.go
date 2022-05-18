package appcontext

import (
	"context"
	"github.com/sirupsen/logrus"
)

type key string

const (
	loggerKey key = "logger"
)

func SetLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey)
	if target, ok := logger.(*logrus.Entry); ok {
		return target
	} else {
		panic("cannot get logger from Context")
	}
}
