package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelDocGroups(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "DELETE", "/doc-groups/:id", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				ID uint `uri:"id"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Delete(&utils.DocGroup{ID: r.ID}); result.RowsAffected == 0 {
					return 400, Res("不存在这个文档组", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("删除文档组失败", nil)
				} else {
					return 200, Res("删除文档组成功", nil)
				}
			},
		),
	)
}
