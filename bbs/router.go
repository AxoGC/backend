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

	cfg := HandlerConfig{
		Env: config,
		Config: p.Config{
			DB:            db,
			JWTKey:        config.JWTKey,
			JWTExpiry:     24 * time.Hour,
			UserTableName: "users",
			AdminColName:  "admin",
			Resper:        &utils.Resper{},
		},
	}

	r.Use(utils.LogMidWare(db))

	utils.RegisterHandlers(r, &cfg,
		ListForumGroups,
		GetForumGroups,
		AddForumGroups,
		EditForumGroups,
		DelForumGroups,
		GetForums,
		AddForums,
		EditForums,
		DelForums,
		GetPosts,
		AddPosts,
		EditPosts,
		DelPosts,
		GetRecommends,
	)

	r.GET("/routes", utils.GetRoutes(r))

	return r
}
