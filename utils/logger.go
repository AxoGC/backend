package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func HandleQueryError(c *gin.Context, err error, msgNotFound, msgError string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, Resp{msgNotFound, nil})
	} else if err != nil {
		c.JSON(500, Resp{msgError, nil})
		c.Error(err)
	}
}
