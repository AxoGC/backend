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

	cfg := HandlerConfig{
		Config: &p.Config{
			DB:            db,
			JWTKey:        config.JWTKey,
			JWTExpiry:     24 * time.Hour,
			UserTableName: "users",
			AdminColName:  "admin",
			Resper:        &utils.Resper{},
		},
		Env: config,
	}

	r.Use(utils.CorsMidWare)
	r.Use(utils.LogMidWare(db))

	utils.RegisterHandlers(r, &cfg,
		ListGuilds,
		GetGuilds,
		AddGuilds,
	)

	r.Use(utils.GetRoutes(r))
	r.Use(utils.LogMidWare(db))

	return r
}
