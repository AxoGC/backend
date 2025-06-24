package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB, rdb *redis.Client) *gin.Engine {
	r := gin.Default()

	cfg := HandlerConfig{}

	utils.RegisterHandlers(r, &cfg,
		ChangePassword,
		GetMyinfo,
		GetUsers,
		Login,
		ResetPassword,
		SetAvatar,
		SetCover,
		SetEmail,
		SetName,
		SetProfile,
		Signup,
	)

	return r
}
