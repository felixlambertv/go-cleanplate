package v1

import (
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	s service.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, l logger.Interface, s service.User) {
	r := &userRoutes{l: l, s: s}

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
	users, err := r.s.GetUsers(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}
	ctx.JSON(http.StatusOK, users)
}
