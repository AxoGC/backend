package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	cfg := p.Config{}

	utils.RegisterHandlers(r, &cfg,
		GetReviews,
		AddReviews,
		EditReviews,
		DelReviews,

		GetCarousels,
	)

	return r
}
