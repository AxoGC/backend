package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/change-password", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				OldPassword string `json:"oldPassword" binding:"required,min=8"`
				NewPassword string `json:"newPassword" binding:"required,min=8"`
			}) (int, *Resp) {

				if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(r.OldPassword)); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					return 400, Res("密码不正确", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("验证密码失败", nil)
				}

				password, err := bcrypt.GenerateFromPassword([]byte(r.NewPassword), bcrypt.DefaultCost)
				if err != nil {
					c.Error(err)
					return 500, Res("密码加密失败", nil)
				}

				if err := cfg.DB.Where(u).Update("password", string(password)).Error; err != nil {
					c.Error(err)
					return 500, Res("密码更新失败", nil)
				}

				return 200, Res("密码更新成功", nil)
			},
		),
	}
}
