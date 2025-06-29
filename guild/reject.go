package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func Reject(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/reject", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON}, nil,
		WithGuildRolesAuth(
			[]utils.DictID{Admin, Owner},
			func(c *gin.Context, u *utils.User, r *struct {
				IDs []uint `json:"ids"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Where(
					"id IN ? AND guild_id IN ? AND guild_role = ?", r.IDs, u.GuildID, Applicant,
				).Update("guild_role", None).Error; err != nil {
					c.Error(err)
					return 500, Res("拒绝申请失败", nil)
				}

				return 200, Res("拒绝申请成功", nil)
			},
		),
	)
}
