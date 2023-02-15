package app

import (
	"fmt"
	"github.com/felixlambertv/go-cleanplate/config"
	v1 "github.com/felixlambertv/go-cleanplate/internal/controller/http/v1"
	"github.com/felixlambertv/go-cleanplate/pkg/httpserver"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l)
	httpServer := httpserver.NewServer(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app run: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("%w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("%w", err))
	}
}
