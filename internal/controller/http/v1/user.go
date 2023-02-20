package v1

import (
	"net/http"
	"strings"

	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/middleware"
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, errorRes{
			Message: "request not valid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	txHandle := ctx.MustGet("db_trx").(*gorm.DB)
	user, err := r.s.WithTrx(txHandle).CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			errorResponse(ctx, http.StatusConflict, errorRes{
				Message: "Duplicate email",
				Debug:   err,
				Errors:  err.Error(),
			})
		} else {
			errorResponse(ctx, http.StatusInternalServerError, errorRes{
				Message: "Something went wrong",
				Debug:   err,
				Errors:  err.Error(),
			})
		}
		return
	}

	successResponse(ctx, http.StatusOK, successRes{
		Message: "Success Creating new user",
		Data:    user,
	})
}

func (r *userRoutes) getUser(ctx *gin.Context) {
	users, err := r.s.GetUsers()
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, errorRes{
			Message: "User not found",
			Debug:   nil,
			Errors:  nil,
		})
		return
	}
	successResponse(ctx, http.StatusOK, successRes{
		Message: "Success Get Users",
		Data:    users,
	})
}
