package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDocGroups(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "GET", "/doc-groups/:doc-group", p.Preload(
		cfg, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			DocGroup string `uri:"doc-group"`
		}) (int, *utils.Resp) {

			type DocGroup struct {
				Label string `json:"label"`
				Slug  string `json:"slug"`
				Sort  string `json:"sort"`
			}

			var docGroup DocGroup
			if err := cfg.DB.Take(&docGroup, "slug = ?", r.DocGroup).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到文档组", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找文档组失败", nil)
			}

			return 200, Res("", &docGroup)
		},
	)
}
