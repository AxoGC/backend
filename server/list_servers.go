package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func ListServers(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/servers", func(c *gin.Context) {

		type Data struct {
			Slug   string `json:"slug"`
			Label  string `json:"label"`
			Port   uint16 `json:"port"`
			Game   string `json:"game"`
			Online *int64 `json:"online"`
		}

		var datas []Data
		var servers []utils.Server
		if err := cfg.DB.Preload("Game").Find(&servers).Error; err != nil {
			c.JSON(500, Res("获取服务器列表失败", nil))
			return
		}

		for _, srv := range servers {
			online, err := srv.GetOnlineCount(cfg.Env.ServerHost)
			if err != nil {
				c.Error(err)
			}
			datas = append(datas, Data{
				Slug:   srv.Slug,
				Label:  srv.Label,
				Port:   srv.Port,
				Game:   srv.Game.Label,
				Online: online,
			})
		}

		c.JSON(200, Res("", datas))
	}
}
