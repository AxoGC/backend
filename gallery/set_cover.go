package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
)

func RegisterCover(r *gin.Engine, cfg *p.Config, config *Config) {
	r.POST("/cover", p.Preload(
		cfg, &p.Option{Bind: p.Other, Permission: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `form:"slug" binding:"required,min=3,alphanum"`
		}) {

			utils.UploadImageMidWare(config.FilePath, "album-cover/", r.Slug)(c)
		},
	))
}
