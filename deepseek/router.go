package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(utils.LogMidWare(db))
	return r
}
