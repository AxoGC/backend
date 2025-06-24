package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditGuilds(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "PATCH", "/edit", []gin.HandlerFunc{
		CheckRoleMidWare(cfg.Config, Admin, Owner),
		p.Preload(
			cfg.Config, &p.Option{Bind: p.JSON, Permission: p.Login}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Name    string `json:"name" binding:"required,min=3"`
				Path    string `json:"path" binding:"required,min=3,alphanum"`
				Profile string `json:"profile"`
				Notice  string `json:"notice"`
			}) (int, *Resp) {

				if err := cfg.DB.Where(&utils.Guild{ID: *u.GuildID}).Updates(r).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("此公会名或标识已被使用", nil)
				} else if err != nil {
					return 500, Res("更新公会信息失败", nil)
				}

				return 200, Res("更新公会信息成功", nil)
			},
		),
	}
}
