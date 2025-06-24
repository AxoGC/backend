package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelForumGroups(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "DELETE", "/forum-groups/:id", []gin.HandlerFunc{
		p.Preload(
			&cfg.Config, &p.Option{Permission: p.Admin, Bind: p.Uri}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				ID uint `uri:"id" binding:"required"`
			}) (int, *Resp) {

				if result := cfg.DB.Delete(&utils.ForumGroup{ID: r.ID}); result.RowsAffected == 0 {
					return 400, Res("不存在此论坛组", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("删除论坛组失败", nil)
				} else {
					return 200, Res("删除论坛组成功", nil)
				}
			},
		),
	}
}
