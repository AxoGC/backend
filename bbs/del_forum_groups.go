package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelForumGroups(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "DELETE", "/forum-groups/:id", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.URI, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				ID uint `uri:"id"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Delete(&utils.ForumGroup{ID: r.ID}); result.RowsAffected == 0 {
					return 404, Res("不存在此论坛组", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("删除论坛组失败", nil)
				} else {
					return 200, Res("删除论坛组成功", nil)
				}
			},
		),
	)
}
