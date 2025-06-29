package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelPosts(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "DELETE", "/posts/:slug", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.URI, Preloads: []string{"UserRoles"}}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
		}) (int, *utils.Resp) {

			var post utils.Post
			if err := cfg.DB.Take(&post, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的帖子", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找帖子失败", nil)
			} else if post.UserID != u.ID && !u.HasAnyRole(utils.Admin, utils.BBSAdmin) {
				return 403, Res("你没有权限删除该内容", nil)
			}

			if err := cfg.DB.Delete(new(utils.Post), "slug = ?", r.Slug).Error; err != nil {
				c.Error(err)
				return 500, Res("帖子删除失败", nil)
			} else {
				return 200, Res("帖子删除成功", nil)
			}
		},
	)
}
