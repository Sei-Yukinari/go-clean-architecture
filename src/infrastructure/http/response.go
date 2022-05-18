package http

import "github.com/gin-gonic/gin"

type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"apperror message"`
}

func NewError(c *gin.Context, status int, err error) {
	er := Error{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}

type Response struct {
	Message string      `json:"message" example:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(c *gin.Context, status int, message string, data interface{}) {
	res := Response{
		Message: message,
		Data:    data,
	}
	c.JSON(status, res)
}
