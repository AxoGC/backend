package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetForumGroups(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/forum-groups/:forumGroup", p.Preload(
		&cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ForumGroup string `uri:"forumGroup"`
		}) (int, *utils.Resp) {

			var forumGroup utils.ForumGroup
			if err := cfg.DB.Take(&forumGroup, "slug = ?", r.ForumGroup).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的论坛组", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找论坛组失败", nil)
			}

			return 200, Res("", &forumGroup)
		},
	)
}
