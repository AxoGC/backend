package main

import (
	"errors"
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGuilds(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/:path", p.Preload(
		cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
		}) (int, *utils.Resp) {

			var guild struct {
				CreatedAt time.Time `json:"createdAt"`
				Name      string    `json:"name"`
				Path      string    `json:"path"`
				Count     uint      `json:"count"`
				Profile   string    `json:"profile"`
				Money     uint      `json:"money"`
				Users     []struct {
					ID        uint   `json:"id"`
					Name      string `json:"name"`
					GuildID   *uint  `json:"guildId"`
					GuildRole uint   `json:"guildRole"`
				} `json:"users"`
			}
			if err := cfg.DB.Where(&utils.Guild{Slug: r.Slug}).Take(&guild).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 400, Res("不存在这个公会", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找公会失败", nil)
			}

			return 200, Res("", &guild)
		},
	)
}
