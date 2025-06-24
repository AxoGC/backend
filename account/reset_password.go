package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ResetPassword(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/reset-password", []gin.HandlerFunc{
		utils.AuthEmailMidWare(cfg.rdb), p.Preload(
			cfg.Config, &p.Option{Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Email    string `json:"email" binding:"required,email"`
				Password string `json:"password" binding:"required,min=8"`
			}) (int, *Resp) {

				password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
				if err != nil {
					c.Error(err)
					return 500, Res("密码加密失败", nil)
				}

				if err := cfg.DB.Where(&utils.User{Email: r.Email}).Update("password", string(password)).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 400, Res("此邮箱尚未注册", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查询用户信息失败", nil)
				}

				return 200, Res("密码更新成功", nil)
			},
		),
	}
}
