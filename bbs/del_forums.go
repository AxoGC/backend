package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelForums(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "DELETE", "/forums/:slug", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.URI, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Delete(new(utils.Forum), "slug = ?", r.Slug); result.RowsAffected == 0 {
					return 404, Res("找不到对应的论坛", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("论坛删除失败", nil)
				} else {
					return 200, Res("论坛删除成功", nil)
				}
			},
		),
	)
}
