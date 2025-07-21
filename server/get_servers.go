package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetServers(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/servers/:slug", p.Preload(
		cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
		}) (int, *utils.Resp) {

			var srv utils.Server
			if err := cfg.DB.Preload("Game").Preload("DocGroup").Preload("Forum").Take(
				&srv, "slug = ?", r.Slug,
			).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到这个服务器", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找服务器失败", nil)
			}

			var status bool

			online, err := utils.GetOnlineCount(srv.GameID, srv.Port, cfg.Env.ServerHost)
			if err != nil && err != utils.ErrNotSupportedGame {
				c.Error(err)
			}

			return 200, Res("", gin.H{
				"label":       srv.Label,
				"port":        srv.Port,
				"description": srv.Description,
				"game":        srv.Game.Label,
				"backup":      srv.BackupEnable,
				"status":      status,
				"online":      online,
				"docGroup":    srv.DocGroup,
				"forum":       srv.Forum,
			})
		},
	)
}
