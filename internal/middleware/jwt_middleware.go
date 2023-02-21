package middleware

import (
	"net/http"

	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(allowedLevel ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		extractedToken := utils.ExtractToken(ctx)
		_, err := utils.ParseToken(extractedToken)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
				Message: "Invalid token",
				Debug:   err,
				Errors:  err.Error(),
			})
			ctx.Abort()
			return
		}

		// uLevel, errULevel := utils.ExtractUserLevel(parsedToken)
		// if errULevel != nil {
		// 	utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
		// 		Message: "Cannot extract user level",
		// 		Debug: errULevel,
		// 		Errors: errULevel.Error(),
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()
	}
}
