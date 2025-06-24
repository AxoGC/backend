package main

import (
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func SetProfile(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/profile", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				IsMale   bool      `json:"isMale" binding:"required"`
				Profile  string    `json:"profile" binding:"required"`
				Birthday time.Time `json:"birthday" binding:"required"`
				Location string    `json:"location" binding:"required"`
			}) (int, *Resp) {

				if err := cfg.DB.Where(u).Updates(r).Error; err != nil {
					c.Error(err)
					return 500, Res("用户信息更新失败", nil)
				}

				return 200, Res("用户信息更新成功", nil)
			},
		),
	}
}
