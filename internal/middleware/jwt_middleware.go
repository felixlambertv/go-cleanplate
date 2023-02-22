package middleware

import (
	"net/http"

	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func JWTAuthMiddleware(allowedLevel ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		extractedToken := utils.ExtractToken(ctx)
		parsedToken, err := utils.ParseToken(extractedToken)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
				Message: "Invalid token",
				Debug:   err,
				Errors:  err.Error(),
			})
			ctx.Abort()
			return
		}

		uLevel, errULevel := utils.ExtractUserLevel(parsedToken)
		if errULevel != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
				Message: "Cannot extract user level",
				Debug:   errULevel,
				Errors:  errULevel.Error(),
			})
			ctx.Abort()
			return
		}

		if !slices.Contains(allowedLevel, uLevel) {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
				Message: "You're not authorized to access this",
				Debug:   nil,
				Errors:  "",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
