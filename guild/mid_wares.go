package main

import (
	"slices"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

/*
func CheckJoinMidWare(cfg *p.Config, join bool) gin.HandlerFunc {
	return p.Preload(
		cfg, &p.Option{Permission: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) {
			if join && u.GuildID == nil {
				c.AbortWithStatusJSON(400, Resp{"你尚未加入公会", nil})
				return
			}
			if !join && u.GuildID != nil {
				c.AbortWithStatusJSON(400, Resp{"你已经加入公会了", nil})
				return
			}
		},
	)
}
*/

func CheckRoleMidWare(cfg *p.Config, roles ...utils.GuildRole) gin.HandlerFunc {
	return p.Preload(
		cfg, &p.Option{Permission: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *Resp) {
			if !slices.Contains(roles, u.GuildRole) {
				return 400, Res("你的公会角色不满足", nil)
			}
			return 0, nil
		},
	)
}

/*

func CheckGuildMidWare(cfg *p.Config, joined bool, roles ...utils.GuildRole) gin.HandlerFunc {
	return p.Preload(
		cfg, &p.Option{Permission: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) {},
	)
}
*/
