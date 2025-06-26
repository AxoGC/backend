package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetCover(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/forums/:slug/cover", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.Other, Preloads: []string{"Roles", "Roles.Role"}}, nil,
		utils.WithRolesAuth(
			[]utils.Role{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug" binding:"required,min=3,alphanum"`
			}) (int, *Resp) {

				if err := cfg.DB.Take(new(utils.Forum), "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 400, Res("不存在这个论坛", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("获取论坛失败", nil)
				}
				return utils.UploadImageMidWare(cfg.Env.FilePath, "forum-cover/", r.Slug)(c)
			},
		),
	)
}
