package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func SetCover(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/albums/:slug/cover", p.Preload(
		cfg.Config, &p.Option{Bind: p.Other, Login: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug" binding:"required,min=3,alphanum"`
		}) (int, *utils.Resp) {

			return utils.UploadImageMidWare(cfg.Env.FilePath, "album-cover/", r.Slug)(c)
		},
	)
}
