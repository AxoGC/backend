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

	r.Use(utils.CorsMidWare)

	cfg := &p.Config{
		DB:            db,
		JWTKey:        config.JWTKey,
		JWTExpiry:     time.Hour * 24,
		UserTableName: "users",
		AdminColName:  "admin",
		Resper:        &utils.Resper{},
	}

	r.Use(utils.LogMidWare(cfg.DB))

	utils.RegisterHandlers(r, cfg,
		ListDocGroups,
		AddDocGroups,
		GetDocGroups,
		EditDocGroups,
		DelDocGroups,
		GetDocs,
		AddDocs,
		EditDocs,
		DelDocs,
	)

	r.GET("routes", utils.GetRoutes(r))

	return r
}
