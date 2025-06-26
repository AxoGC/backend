package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/signup", utils.WithEmailAuth(cfg.rdb, p.Preload(
		cfg.Config, &p.Option{Bind: p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Name     string `json:"name" binding:"required,min=3,max=32"`
			Password string `json:"password" binding:"required,min=8"`
			Email    string `json:"email" binding:"required,email"`
		}) (int, *Resp) {

			if err := cfg.DB.Take(new(utils.User), "name = ?", r.Name).Error; err == nil {
				return 400, Res("此用户名已被注册", nil)
			}

			if err := cfg.DB.Take(new(utils.User), "email = ?", r.Email).Error; err == nil {
				return 400, Res("此邮箱已被使用", nil)
			}

			password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
			if err != nil {
				c.Error(err)
				return 500, Res("密码加密失败", nil)
			}

			user := utils.User{
				Name:     r.Name,
				Password: string(password),
				Email:    r.Email,
			}

			if err := cfg.DB.Create(&user).Error; err != nil {
				c.Error(err)
				return 500, Res("创建用户失败", nil)
			}

			token, err := p.GetJwt(user.ID, cfg.JWTKey, cfg.JWTExpiry)
			if err != nil {
				c.Error(err)
				return 500, Res("用户凭证生成失败", nil)
			}

			return 200, Res("用户创建成功", token)
		},
	))
}
