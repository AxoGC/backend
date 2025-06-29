package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddPosts(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/forums/:forumSlug/posts", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ForumSlug string `uri:"forumSlug"`
			Slug      string `json:"slug"`
			Title     string `json:"title"`
			Content   string `json:"content"`
			Markdown  bool   `json:"markdown"`
		}) (int, *utils.Resp) {

			var forum utils.Forum
			if err := cfg.DB.Take(&forum, "slug = ?", r.ForumSlug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的论坛", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找论坛失败", nil)
			}

			if err := cfg.DB.Model(new(utils.Post)).Create(map[string]any{
				"slug":     r.Slug,
				"title":    r.Title,
				"content":  r.Content,
				"markdown": r.Markdown,
				"forum_id": forum.ID,
				"user_id":  u.ID,
			}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
				return 409, Res("对应的标题或标识已被占用", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("帖子创建失败", nil)
			} else {
				return 201, Res("帖子创建成功", nil)
			}
		},
	)
}
