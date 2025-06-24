package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddPosts(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/forums/:forumSlug/posts", []gin.HandlerFunc{
		p.Preload(
			&cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				ForumSlug string `uri:"forumSlug" binding:"required"`
				Slug      string `json:"slug" binding:"required,min=3"`
				Title     string `json:"title" binding:"required,min=3"`
				Content   string `json:"content"`
				Markdown  bool   `json:"markdown"`
			}) (int, *Resp) {

				var forum utils.Forum
				if err := cfg.DB.Take(&forum, "slug = ?", r.ForumSlug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("找不到对应的论坛", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找论坛失败", nil)
				}

				if err := cfg.DB.Model(new(utils.Post)).Create(r).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("对应的标题或标识已被占用", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("帖子创建失败", nil)
				} else {
					return 201, Res("帖子创建成功", nil)
				}
			},
		),
	}
}
