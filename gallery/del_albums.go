package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func AddAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/albums", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
			}) (int, *Resp) {
			},
		),
	}
}
