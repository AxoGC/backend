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

func GetRecommends(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/recommends", p.Preload(
		&cfg.Config, &p.Option{Bind: p.Query}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Option string `form:"option"`
			Count  int    `form:"count"`
			Forum  string `form:"forum"`
		}) (int, *utils.Resp) {

			if r.Count < 0 || r.Count > 100 {
				return 400, Res("请输入0 ~ 100的数字", nil)
			}

			type Forum struct {
				ID    uint   `json:"id"`
				Title string `json:"title"`
				Slug  string `json:"slug"`
			}

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type Post struct {
				ID          uint      `json:"id"`
				UpdatedAt   time.Time `json:"updatedAt"`
				Title       string    `json:"title"`
				Slug        string    `json:"slug"`
				ForumID     uint      `json:"forumId"`
				Forum       Forum     `json:"forum"`
				UserID      uint      `json:"userId"`
				User        User      `json:"user"`
				ReviewCount uint      `json:"reviewCount"`
			}

			var posts []Post
			query := cfg.DB.Model(new(utils.Post)).Preload("Forum",
				s.Model(new(utils.Forum)),
			).Preload("User",
				s.Model(new(utils.User)),
			).Limit(r.Count)
			switch r.Option {
			case "popular":
				query = query.Order("review_count DESC")
			case "latest":
				query = query.Order("created_at DESC")
			default:
				return 400, Res("不存在此选项", nil)
			}

			if r.Forum != "" {
				var forumId uint
				if err := cfg.DB.Model(new(utils.Forum)).Where(
					"slug = ?", r.Forum,
				).Pluck("id", &forumId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("不存在这个论坛", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("论坛查找失败", nil)
				}
				query = query.Where("forum_id = ?", forumId)
			}

			if err := query.Find(&posts).Error; err != nil {
				c.Error(err)
				return 500, Res("查询帖子失败", nil)
			}

			return 200, Res("", posts)
		},
	)
}
