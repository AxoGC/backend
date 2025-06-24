package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelGuilds(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/disband", []gin.HandlerFunc{
		CheckRoleMidWare(cfg.Config, Owner),
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login}, nil,
			func(c *gin.Context, u *utils.User, r *struct{}) (int, *Resp) {

				if err := cfg.DB.Delete(utils.Guild{ID: *u.GuildID}).Error; err != nil {
					c.Error(err)
					return 500, Res("公会解散失败", nil)
				}

				return 200, Res("公会解散成功", nil)
			},
		),
	}
}
