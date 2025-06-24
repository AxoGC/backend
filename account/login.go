package main

import (
	"errors"
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/login", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Name     string `json:"name" binding:"required,min=3,max=32"`
				Password string `json:"password" binding:"required,min=8"`
			}) (int, *Resp) {

				var user utils.User
				if err := cfg.DB.First(&user, "name = ?", r.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("你尚未注册", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("获取用户信息失败", nil)
				}

				if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					return 400, Res("密码不正确", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("密码验证失败", nil)
				}

				token, err := p.GetJwt(user.ID, cfg.JWTKey, 7*24*time.Hour)
				if err != nil {
					c.Error(err)
					return 500, Res("生成用户凭证失败", nil)
				}

				return 200, Res("", token)
			},
		),
	}
}
