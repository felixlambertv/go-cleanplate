package v1

import "github.com/gin-gonic/gin"

type errorRes struct {
	Message string `json:"message"`
	Debug   error  `json:"debug,omitempty"`
	Errors  any    `json:"errors"`
}

type successRes struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func errorResponse(c *gin.Context, code int, res errorRes) {
	c.JSON(code, res)
}

func successResponse(c *gin.Context, code int, res successRes) {
	c.JSON(code, res)
}
