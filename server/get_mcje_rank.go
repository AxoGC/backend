package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func GetMCJERank(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "GET", "/:slug/rank", []gin.HandlerFunc{
		p.Preload(cfg.Config, &p.Option{Bind: p.URI}, nil, func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
		}) (int, *utils.Resp) {
			return 200, Res("", nil)
		}),
	}
}
