package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func AddAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/albums", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Slug    string `json:"slug"`
				Label   string `json:"label"`
				Profile string `json:"profile"`
				Guild   bool   `json:"guild"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Create(&utils.Album{
					UserID:  u.ID,
					Slug:    r.Slug,
					Label:   r.Label,
					Profile: r.Profile,
					GuildID: lo.Ternary(r.Guild, u.GuildID, nil),
				}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 400, Res("已存在同名相册", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("相册创建失败", nil)
				}

				return 200, Res("相册创建成功", nil)
			},
		),
	}
}
