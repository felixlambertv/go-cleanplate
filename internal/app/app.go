package app

import (
	"fmt"
	"github.com/felixlambertv/go-cleanplate/config"
	v1 "github.com/felixlambertv/go-cleanplate/internal/controller/http/v1"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/felixlambertv/go-cleanplate/pkg/httpserver"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	db, err := gorm.Open(postgres.Open(cfg.PG.GetDbConnectionUrl()))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres: %w", err))
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - migrate: %w", err))
	}

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, db)
	httpServer := httpserver.NewServer(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app run: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("%w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("%w", err))
	}
}
