package main

import (
	"github.com/gin-gonic/gin"
)

func GetRouter(cfg *HandlerConfig) *gin.Engine {
	r := gin.Default()
	return r
}
