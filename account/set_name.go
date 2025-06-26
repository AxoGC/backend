package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func SetName(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/name", p.Preload(
		cfg.Config, &p.Option{Bind: p.JSON, Login: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Name string `json:"name" binding:"required,min=3,max=32"`
		}) (int, *Resp) {

			if err := cfg.DB.Take(new(utils.User), "name = ?", r.Name).Error; err == nil {
				return 400, Res("此名称已有人使用", nil)
			}

			if err := cfg.DB.Where(u).Update("name", r.Name).Error; err != nil {
				c.Error(err)
				return 500, Res("用户名更新失败", nil)
			}

			return 200, Res("用户名更新成功", nil)
		},
	)
}
