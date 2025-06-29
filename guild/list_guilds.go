package main

import (
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func ListGuilds(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/list", p.Preload(
		cfg.Config, &p.Option{}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *utils.Resp) {

			var guilds []struct {
				ID        uint      `json:"id"`
				CreatedAt time.Time `json:"createdAt"`
				Name      string    `json:"name"`
				Count     uint      `json:"count"`
				Profile   string    `json:"profile"`
			}
			if err := cfg.DB.Model(new(utils.Guild)).Scopes(utils.Paginate(c, nil)).Find(&guilds).Error; err != nil {
				c.Error(err)
				return 500, Res("获取公会列表失败", nil)
			}

			return 200, Res("获取公会列表成功", guilds)
		},
	)
}
