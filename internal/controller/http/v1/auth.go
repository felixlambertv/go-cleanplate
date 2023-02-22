package v1

import (
	"net/http"

	"github.com/felixlambertv/go-cleanplate/config"
	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/middleware"
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authRoutes struct {
	s   service.IAuthService
	l   logger.Interface
	cfg *config.Config
}

func newAuthRoutes(handler *gin.RouterGroup, l logger.Interface, db *gorm.DB, s service.IAuthService, cfg *config.Config) {
	r := &authRoutes{l: l, s: s, cfg: cfg}

	h := handler.Group("auth")
	{
		h.POST("login", r.login, middleware.JWTAuthMiddleware(cfg))
		h.POST("register", r.register)
	}
}

func (r *authRoutes) login(ctx *gin.Context) {
	var req request.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorRes{
			Message: "request not valid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	user, _, err := r.s.Login(req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
			Message: "User not found",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, utils.SuccessRes{
		Message: "Login Successful",
		Data:    user,
	})
}

func (r *authRoutes) register(ctx *gin.Context) {
	var req request.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorRes{
			Message: "request not valid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	user, _, err := r.s.Login(req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
			Message: "User not found",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, utils.SuccessRes{
		Message: "Login Successful",
		Data:    user,
	})
}
