package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/users/:name", p.Preload(
		cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Name string `uri:"name" binding:"required"`
		}) (int, *Resp) {

			var user utils.User
			if err := cfg.DB.Where(utils.User{Name: r.Name}).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("不存在该用户", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查询用户失败", nil)
			}

			return 200, Res("", gin.H{
				"id":   user.ID,
				"name": user.Name,
			})
		},
	)
}
