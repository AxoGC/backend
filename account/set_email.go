package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func SetEmail(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/email", []gin.HandlerFunc{
		utils.AuthEmailMidWare(cfg.rdb), p.Preload(
			cfg.Config, &p.Option{Bind: p.JSON, Permission: p.Login}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Email string `json:"email" binding:"required,email"`
			}) (int, *Resp) {

				if err := cfg.DB.Take(new(utils.User), "email = ?", r.Email).Error; err == nil {
					return 400, Res("此邮箱已经有人使用", nil)
				}

				if err := cfg.DB.Where(u).Update("email", r.Email).Error; err != nil {
					c.Error(err)
					return 500, Res("邮箱更新失败", nil)
				}

				return 200, Res("邮箱更新成功", nil)
			},
		),
	}
}
