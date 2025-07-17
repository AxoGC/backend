package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDocs(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "POST", "/doc-groups/:slug/docs", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				DocGroupSlug string `uri:"slug"`
				Slug         string `json:"slug"`
				Title        string `json:"title"`
				Content      string `json:"content"`
				Sort         int    `json:"sort"`
			}) (int, *utils.Resp) {

				var docGroup utils.DocGroup
				if err := cfg.DB.Take(&docGroup, "slug = ?", r.DocGroupSlug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("不存在对应的文档组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找文档组失败", nil)
				}

				if err := cfg.DB.Model(new(utils.Doc)).Create(&utils.Doc{
					DocGroupID: docGroup.ID,
					Slug:       r.Slug,
					Title:      r.Title,
					Content:    r.Content,
					Sort:       r.Sort,
					UserID:     u.ID,
				}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("已存在相同名称或标识的文档", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("创建文档失败", nil)
				} else {
					return 201, Res("文档创建成功", nil)
				}
			},
		),
	)
}
