package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDocGroups(cfg *p.Config) (string, string, []gin.HandlerFunc) {
	return "POST", "/doc-groups", []gin.HandlerFunc{
		p.Preload(
			cfg, &p.Option{Permission: p.Admin, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Label string `json:"label" binding:"required"`
				Sort  int    `json:"sort"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Model(new(utils.DocGroup)).Create(r).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 400, Res("已存在同名文档组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("创建文档组失败", nil)
				} else {
					return 200, Res("创建文档组成功", nil)
				}
			},
		),
	}
}
