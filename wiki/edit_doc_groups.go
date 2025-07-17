package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditDocGroups(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "PUT", "/doc-groups/:slug", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				OldSlug string `uri:"slug"`
				Slug    string `json:"slug"`
				Label   string `json:"label"`
				Sort    int    `json:"sort"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Model(new(utils.DocGroup)).Where("slug = ?", r.OldSlug).Updates(map[string]any{
					"label": r.Label,
					"slug":  r.Slug,
					"sort":  r.Sort,
				}); errors.Is(result.Error, gorm.ErrDuplicatedKey) {
					return 400, Res("此名称或标识已被其他文档组占用", nil)
				} else if result.RowsAffected == 0 {
					return 400, Res("没有对应的文档组", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("更新文档组失败", nil)
				} else {
					return 200, Res("更新文档组成功", nil)
				}
			},
		),
	)
}
