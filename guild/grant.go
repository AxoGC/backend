package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func Grant(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/grant", []gin.HandlerFunc{
		CheckRoleMidWare(cfg.Config, Owner), p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				IDs []uint `json:"ids"`
			}) (int, *Resp) {

				if err := cfg.DB.Where(
					"id IN ? AND guild_id IN ? AND guild_role = ?", r.IDs, u.GuildID, Member,
				).Update("guild_role", Admin).Error; err != nil {
					return 500, Res("提拔为管理员失败", nil)
				}

				return 200, Res("提拔为管理员成功", nil)
			},
		),
	}
}
