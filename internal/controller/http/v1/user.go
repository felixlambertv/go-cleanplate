package v1

import (
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, l logger.Interface) {
	r := &userRoutes{l: l}

	h := handler.Group("users")
	{
		h.GET("", r.getUser)
		h.GET("aa", r.getUser)
	}
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *userRoutes) getUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userResponse{Name: "Nice", Email: "nice@nice.com"})
}
