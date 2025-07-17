package main

import (
	"errors"
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPosts(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/posts/:slug", p.Preload(
		&cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
		}) (int, *utils.Resp) {

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type Forum struct {
				ID    uint   `json:"id"`
				Title string `json:"title"`
			}

			type Review struct {
				ID             uint      `json:"id"`
				UpdatedAt      time.Time `json:"updatedAt"`
				Content        string    `json:"content"`
				Attitude       *bool     `json:"attitude"`
				UserID         uint      `json:"userId"`
				ReviewableID   uint      `json:"reviewableId"`
				ReviewableType string    `json:"reviewableType"`
				User           User      `json:"user"`
				ReviewCount    uint      `json:"reviewCount"`
			}

			type Post struct {
				ID          uint      `json:"id"`
				CreatedAt   time.Time `json:"createdAt"`
				UpdatedAt   time.Time `json:"updatedAt"`
				Pinned      bool      `json:"pinned"`
				Title       string    `json:"title"`
				ForumID     uint      `json:"forumId"`
				Content     string    `json:"content"`
				Markdown    bool      `json:"markdown"`
				UserID      uint      `json:"userId"`
				ReviewCount uint      `json:"reviewCount"`
				User        User      `json:"user"`
				Forum       Forum     `json:"forum"`
				Reviews     []Review  `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:posts"`
			}

			var post Post

			if err := cfg.DB.Preload(
				"Reviews", s.Preload("User"),
			).Preload("User").Preload("Forum").Take(&post, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("不存在这个帖子", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("获取帖子数据失败", nil)
			} else {
				return 200, Res("", &post)
			}
		},
	)
}
