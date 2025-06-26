package utils

import (
	"github.com/gin-gonic/gin"
)

type Resper struct{}

func (r *Resper) Resp(c *gin.Context, status int, message string, err error) {
	if err != nil {
		c.Error(err)
	}
	c.JSON(status, Resp{message, nil})
}

type Resp struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Res(message string, data any) *Resp {
	return &Resp{Message: message, Data: data}
}
