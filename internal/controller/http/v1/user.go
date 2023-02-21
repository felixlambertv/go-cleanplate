package v1

import (
	"net/http"
	"strings"

	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/middleware"
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
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
		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorRes{
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
			utils.ErrorResponse(ctx, http.StatusConflict, utils.ErrorRes{
				Message: "Duplicate email",
				Debug:   err,
				Errors:  err.Error(),
			})
		} else {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, utils.ErrorRes{
				Message: "Something went wrong",
				Debug:   err,
				Errors:  err.Error(),
			})
		}
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, utils.SuccessRes{
		Message: "Success Creating new user",
		Data:    user,
	})
}

func (r *userRoutes) getUser(ctx *gin.Context) {
	users, err := r.s.GetUsers()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, utils.ErrorRes{
			Message: "User not found",
			Debug:   nil,
			Errors:  nil,
		})
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, utils.SuccessRes{
		Message: "Success Get Users",
		Data:    users,
	})
}
