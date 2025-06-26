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
			ID uint `uri:"id" binding:"required"`
		}) (int, *Resp) {

			var post struct {
				CreatedAt   time.Time `json:"createdAt"`
				UpdatedAt   time.Time `json:"updatedAt"`
				Pinned      bool      `json:"pinned"`
				Title       string    `json:"title"`
				ForumID     uint      `json:"forumId"`
				Content     string    `json:"content"`
				Markdown    bool      `json:"markdown"`
				UserID      uint      `json:"userId"`
				ReviewCount uint      `json:"reviewCount"`
				User        struct {
					ID   uint   `json:"id"`
					Name string `json:"name"`
				} `json:"user"`
				Forum struct {
					ID    uint   `json:"id"`
					Title string `json:"title"`
				} `json:"forum"`
				Reviews []struct {
					ID             uint      `json:"id"`
					UpdatedAt      time.Time `json:"updatedAt"`
					Content        string    `json:"content"`
					Attitude       *bool     `json:"attitude"`
					UserID         uint      `json:"userId"`
					ReviewableID   uint      `json:"reviewableId"`
					ReviewableType string    `json:"reviewableType"`
					User           struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
					ReviewCount uint `json:"reviewCount"`
				} `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:posts"`
			}

			if err := cfg.DB.Model(new(utils.Post)).Preload("Reviews",
				s.Model(new(utils.Review)),
				s.Preload("User",
					s.Model(new(utils.User)),
				),
			).Preload("User",
				s.Model(new(utils.User)),
			).Preload("Forum",
				s.Model(new(utils.Forum)),
			).First(&post, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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
