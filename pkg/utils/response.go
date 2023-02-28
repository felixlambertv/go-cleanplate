package utils

import "github.com/gin-gonic/gin"

type ErrorRes struct {
	Message string `json:"message"`
	Debug   error  `json:"debug,omitempty"`
	Errors  any    `json:"errors"`
}

type SuccessRes struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, code int, res ErrorRes) {
	c.JSON(code, res)
}

func SuccessResponse(c *gin.Context, code int, res SuccessRes) {
	c.JSON(code, res)
}
