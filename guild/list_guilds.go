package main

import (
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func ListGuilds(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/guilds", p.Preload(
		cfg.Config, &p.Option{}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *utils.Resp) {

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type UserGuild struct {
				ID                uint                    `json:"id"`
				UserID            uint                    `json:"userId"`
				User              User                    `json:"user"`
				UserGuildStatusID utils.UserGuildStatusID `json:"userGuildStatusId"`
				GuildID           uint                    `json:"guildId"`
			}

			type Guild struct {
				ID         uint        `json:"id"`
				CreatedAt  time.Time   `json:"createdAt"`
				Name       string      `json:"name"`
				Slug       string      `json:"slug"`
				SubTitle   string      `json:"subTitle"`
				UserCount  uint        `json:"userCount"`
				UserGuilds []UserGuild `json:"userGuilds"`
			}

			var guilds []Guild
			if err := cfg.DB.Preload(
				"UserGuilds", "user_guild_status_id = ?", "guild_admin",
			).Preload("UserGuilds.User").Scopes(utils.Paginate(c, nil)).Find(&guilds).Error; err != nil {
				c.Error(err)
				return 500, Res("获取公会列表失败", nil)
			}

			return 200, Res("", guilds)
		},
	)
}
