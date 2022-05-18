package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(handler *gin.Engine) {
	srv := &http.Server{
		Addr:    ":" + config.Conf.App.Port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Errorf("apperror")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.Conf.App.TimeoutToGracefulShutdownMs)*time.Millisecond,
	)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Errorf("Server forced to shutdown:%v\n", err)
	}

	fmt.Println("Server exiting")
}
