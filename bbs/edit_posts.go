package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditPosts(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "PATCH", "/posts/:slug", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"Roles", "Roles.Role"}}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			OldSlug  string `uri:"slug"`
			Slug     string `json:"slug"`
			Title    string `json:"title"`
			ForumID  uint   `json:"forumId"`
			Content  string `json:"content"`
			Markdown bool   `json:"markdown"`
		}) (int, *utils.Resp) {

			var post utils.Post
			if err := cfg.DB.Take(&post, "slug = ?", r.OldSlug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的帖子", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找帖子失败", nil)
			} else if post.UserID != u.ID && !u.HasAnyRole(utils.Admin, utils.BBSAdmin) {
				return 403, Res("你没有权限编辑该内容", nil)
			}

			if result := cfg.DB.Model(new(utils.Post)).Where("slug = ?", r.OldSlug).Updates(map[string]any{
				"slug":     r.Slug,
				"title":    r.Title,
				"forum_id": r.ForumID,
				"content":  r.Content,
				"markdown": r.Markdown,
				"user_id":  u.ID,
			}); errors.Is(result.Error, gorm.ErrCheckConstraintViolated) {
				return 404, Res("找不到对应的论坛", nil)
			} else if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
				return 409, Res("此标识或标题已被其他帖子使用", nil)
			} else if result.Error != nil {
				c.Error(result.Error)
				return 500, Res("帖子更新失败", nil)
			}

			return 200, Res("帖子更新成功", nil)
		},
	)
}
