package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Join(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/join/:slug", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.URI}, nil,
		WithGuildRolesAuth(
			[]utils.DictID{None},
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug"`
			}) (int, *utils.Resp) {

				var guildId uint
				if err := cfg.DB.Where(&utils.Guild{Slug: r.Slug}).Pluck("id", &guildId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("没有对应的公会", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查询公会失败", nil)
				}

				if err := cfg.DB.Where(u).Updates(map[string]any{"guild_id": guildId, "guild_role": 1}).Error; err != nil {
					return 500, Res("申请加入公会失败", nil)
				}

				return 200, Res("申请成功", nil)
			},
		),
	)
}
