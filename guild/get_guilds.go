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
	return "GET", "/guilds/:slug", p.Preload(
		cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
		}) (int, *utils.Resp) {

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type UserGuild struct {
				ID                uint                    `json:"id"`
				UserID            uint                    `json:"userId"`
				User              User                    `json:"user"`
				UserGuildStatusID utils.UserGuildStatusID `json:"userGuildStatusId"`
				UserGuildStatus   utils.UserGuildStatus   `json:"userGuildStatus"`
				GuildID           uint                    `json:"guildId"`
			}

			type Guild struct {
				ID         uint        `json:"id"`
				CreatedAt  time.Time   `json:"createdAt"`
				Name       string      `json:"name"`
				Slug       string      `json:"slug"`
				UserCount  uint        `json:"userCount"`
				SubTitle   string      `json:"subTitle"`
				Profile    string      `json:"profile"`
				Money      uint        `json:"money"`
				UserGuilds []UserGuild `json:"userGuilds"`
			}

			var guild Guild

			if err := cfg.DB.Preload("UserGuilds").Preload(
				"UserGuilds.UserGuildStatus",
			).Preload("UserGuilds.User").Where(
				&utils.Guild{Slug: r.Slug},
			).Take(&guild).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 400, Res("不存在这个公会", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找公会失败", nil)
			}

			return 200, Res("", &guild)
		},
	)
}
