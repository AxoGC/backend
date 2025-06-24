package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	return r
}
