package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelDocGroups(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "DELETE", "/doc-groups/:slug", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug"`
			}) (int, *utils.Resp) {

				var docGroupId uint
				if err := cfg.DB.Model(new(utils.DocGroup)).Where(
					"slug = ?", r.Slug,
				).Pluck("id", &docGroupId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 400, Res("不存在这个文档组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找文档组失败", nil)
				}

				if err := cfg.DB.Delete(new(utils.DocGroup), docGroupId).Error; err != nil {
					c.Error(err)
					return 500, Res("删除文档组失败", nil)
				} else {
					return 200, Res("删除文档组成功", nil)
				}
			},
		),
	)
}
