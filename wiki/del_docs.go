package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelDocs(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "DELETE", "/docs/:slug", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Delete(new(utils.Doc), "slug = ?", r.Slug); result.RowsAffected == 0 {
					return 400, Res("没有对应的数据", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("删除失败", nil)
				} else {
					return 200, Res("删除成功", nil)
				}
			},
		),
	)
}
