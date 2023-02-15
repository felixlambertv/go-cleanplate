package v1

import (
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "oks"})
	})

	h := handler.Group("api/v1")
	{
		newUserRoutes(h, l)
	}
}
