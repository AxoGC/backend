package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func AddAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/albums", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Path    string `json:"path"`
				Title   string `json:"title"`
				Profile string `json:"profile"`
				Guild   bool   `json:"guild"`
			}) (int, *Resp) {

				if err := cfg.DB.Take(new(utils.Album), "path = ?", r.Path).Error; err == nil {
					c.JSON(400, Resp{"已存在同名相册", nil})
				}

				album := utils.Album{
					UserID:  u.ID,
					Path:    r.Path,
					Title:   r.Title,
					Profile: r.Profile,
				}
				if r.Guild {
					if u.GuildID != nil {
						album.GuildID = u.GuildID
					} else {
						c.JSON(400, Resp{"你尚未加入公会", nil})
						return
					}
				}

				if err := cfg.DB.Create(&album).Error; err != nil {
					c.JSON(500, Resp{"相册创建失败", nil})
					return
				}

				c.JSON(200, Resp{"相册创建成功", nil})
			},
		),
	}
}
