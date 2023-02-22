package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/felixlambertv/go-cleanplate/config"
	"github.com/felixlambertv/go-cleanplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

func extractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")

	// Apple already reserved header for Authorization
	// https://developer.apple.com/documentation/foundation/nsurlrequest
	if bearerToken == "" {
		bearerToken = c.Request.Header.Get("X-Authorization")
	}

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func JWTAuthMiddleware(cfg *config.Config, allowedLevel ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		extractedToken := extractToken(ctx)
		parsedToken, err := utils.ParseToken(extractedToken, cfg.App.Secret)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
				Message: "Invalid token",
				Debug:   err,
				Errors:  err.Error(),
			})
			ctx.Abort()
			return
		}

		fmt.Println(parsedToken)

		// uLevel, errULevel := utils.ExtractUserLevel(parsedToken)
		// if errULevel != nil {
		// 	utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
		// 		Message: "Cannot extract user level",
		// 		Debug:   errULevel,
		// 		Errors:  errULevel.Error(),
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		// if !slices.Contains(allowedLevel, uLevel) {
		// 	utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorRes{
		// 		Message: "You're not authorized to access this",
		// 		Debug:   nil,
		// 		Errors:  "",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()
	}
}
