package testutil

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-clean-architecture/src/infrastructure/logger"
	"go-clean-architecture/src/util/appcontext"
)

func SetupContext() context.Context {
	ctx := context.Background()
	l := logger.New()
	return appcontext.SetLogger(ctx, logrus.NewEntry(l))
}
