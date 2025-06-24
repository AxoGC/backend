package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddGuilds(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/create", []gin.HandlerFunc{
		CheckRoleMidWare(cfg.Config, None),
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Name string `json:"name" binding:"required,min=3"`
				Slug string `json:"slug" binding:"required,min=3"`
			}) (int, *Resp) {

				guild := utils.Guild{
					Name:      r.Name,
					Slug:      r.Slug,
					UserCount: 1,
				}

				if err := cfg.DB.Create(&guild).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("此公会名称或标识已被占用", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("公会创建失败", nil)
				}

				if err := cfg.DB.Where(u).Updates(&utils.User{GuildID: &guild.ID, GuildRole: Owner}).Error; err != nil {
					c.Error(err)
					return 500, Res("更新我的公会信息失败", nil)
				}

				return 200, Res("公会创建成功", nil)
			},
		),
	}
}
