package core

import "github.com/gin-gonic/gin"

// NewError example
func NewHTTPError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError example
type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"request completed"`
}
