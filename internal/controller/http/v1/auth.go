package v1

import (
	"encoding/json"
	"net/http"

	"github.com/felixlambertv/go-cleanplate/internal/controller/request"
	"github.com/felixlambertv/go-cleanplate/internal/controller/response"
	"github.com/felixlambertv/go-cleanplate/internal/service"
	"github.com/felixlambertv/go-cleanplate/pkg/logger"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authRoutes struct {
	s service.IAuthService
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, l logger.Interface, db *gorm.DB, s service.IAuthService) {
	r := &authRoutes{l: l, s: s}

	h := handler.Group("auth")
	{
		h.POST("login", r.login)
		h.POST("register", r.register)
	}
}

func (r *authRoutes) login(ctx *gin.Context) {
	var req request.LoginRequest
	var res response.LoginResponse

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorRes{
			Message: "request not valid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	user, token, err := r.s.Login(req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
			Message: "User not found",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	userMarshal, _ := json.Marshal(user)
	json.Unmarshal(userMarshal, &res)
	tokenMarshal, _ := json.Marshal(token)
	json.Unmarshal(tokenMarshal, &res)

	utils.SuccessResponse(ctx, http.StatusOK, utils.SuccessRes{
		Message: "Login Successful",
		Data:    res,
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
