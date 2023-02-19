package v1

import (
	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/middleware"
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type userRoutes struct {
	s service.IUserService
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, l logger.Interface, db *gorm.DB, s service.IUserService) {
	r := &userRoutes{l: l, s: s}

	h := handler.Group("users")
	{
		h.GET("", r.getUser)
		h.POST("", middleware.DbTransactionMiddleware(db), r.createUser)
	}
}

func (r *userRoutes) createUser(ctx *gin.Context) {
	var req request.CreateUserRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, "request not valid")
		return
	}
	txHandle := ctx.MustGet("db_trx").(*gorm.DB)
	user, err := r.s.WithTrx(txHandle).CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			errorResponse(ctx, http.StatusConflict, "Duplicate email", err.Error())
		} else {
			errorResponse(ctx, http.StatusInternalServerError, "Something went wrong", err, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (r *userRoutes) getUser(ctx *gin.Context) {
	users, err := r.s.GetUsers()
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}
	ctx.JSON(http.StatusOK, users)
}
