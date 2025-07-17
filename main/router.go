package main

import (
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	cfg := p.Config{
		DB:            db,
		JWTKey:        config.JWTKey,
		JWTExpiry:     24 * time.Hour,
		UserTableName: "user",
		AdminColName:  "admin",
		Resper:        &utils.Resper{},
	}

	r.Use(utils.CorsMidWare)

	utils.RegisterHandlers(r, &cfg,
		GetReviews,
		AddReviews,
		EditReviews,
		DelReviews,

		GetCarousels,
	)

	r.GET("/routes", utils.GetRoutes(r))

	return r
}
