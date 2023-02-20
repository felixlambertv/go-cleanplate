package v1

import (
	"net/http"

	userR "github.com/felixlambertv/go-cleanplate/internal/repository/user"

	"github.com/felixlambertv/go-cleanplate/internal/service/user"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(handler *gin.Engine, l logger.Interface, db *gorm.DB) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	userRepo := userR.NewUserRepo(db, l)
	userService := user.NewUserService(userRepo)

	handler.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "oks"})
	})

	h := handler.Group("api/v1")
	{
		newUserRoutes(h, l, db, userService)
	}
}
