package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	cfg := HandlerConfig{
		Config: &p.Config{
			DB:     db,
			JWTKey: config.JWTKey,
		},
		Env: config,
	}

	r.Use(utils.CorsMidWare)
	r.Use(utils.LogMidWare(db))

	utils.RegisterHandlers(r, &cfg,
		AddAlbums,
		AddImages,
		DelAlbums,
		DelImages,
		EditAlbums,
		EditImages,
		GetAlbums,
		GetImages,
		ListAlbums,
		SetCover,
		GetGalleryStats,
		GetGalleryCarousel,
		GetGalleryRank,
	)

	r.GET("/routes", utils.GetRoutes(r))

	return r
}
