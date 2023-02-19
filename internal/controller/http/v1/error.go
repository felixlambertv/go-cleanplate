package v1

import "github.com/gin-gonic/gin"

type response struct {
	Message string `json:"message"`
	Errors  []any  `json:"errors"`
}

func errorResponse(c *gin.Context, code int, msg string, errors ...any) {
	var res response
	res.Message = msg
	res.Errors = errors
	c.AbortWithStatusJSON(code, res)
}
