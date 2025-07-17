package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func ListServers(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/servers", func(c *gin.Context) {

		type Game struct {
			ID    utils.GameID `json:"id"`
			Label string       `json:"label"`
		}

		type Server struct {
			Slug   string       `json:"slug"`
			Label  string       `json:"label"`
			Port   uint16       `json:"port"`
			GameID utils.GameID `json:"gameId"`
			Game   Game         `json:"game"`
			Online *int64       `json:"online"`
		}

		var servers []Server
		if err := cfg.DB.Preload("Game").Find(&servers).Error; err != nil {
			c.JSON(500, Res("获取服务器列表失败", nil))
			return
		}

		for idx, srv := range servers {
			online, err := utils.GetOnlineCount(srv.GameID, srv.Port, cfg.Env.ServerHost)
			if err != nil {
				c.Error(err)
			}
			servers[idx].Online = online
		}

		c.JSON(200, Res("", servers))
	}
}
