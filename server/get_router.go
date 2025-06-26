package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB, rdb *redis.Client) *gin.Engine {
	r := gin.Default()
	return r
}
