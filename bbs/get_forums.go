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

func GetForums(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/forums/:slug", p.Preload(
		&cfg.Config, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug" binding:"required"`
		}) (int, *utils.Resp) {

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type Post struct {
				ID          uint      `json:"id"`
				UpdatedAt   time.Time `json:"updatedAt"`
				Pinned      bool      `json:"pinned"`
				Title       string    `json:"title"`
				Slug        string    `json:"slug"`
				ForumID     uint      `json:"forumId"`
				UserID      uint      `json:"userId"`
				User        User      `json:"user"`
				ReviewCount uint      `json:"reviewCount"`
			}

			type Server struct {
				ID    uint   `json:"id"`
				Slug  string `json:"slug"`
				Label string `json:"label"`
			}

			type Forum struct {
				ID        uint    `json:"id"`
				Slug      string  `json:"slug"`
				Title     string  `json:"title"`
				SubTitle  string  `json:"subTitle"`
				Profile   string  `json:"profile"`
				PostCount uint    `json:"postCount"`
				ServerID  *uint   `json:"serverId"`
				Server    *Server `json:"server"`
				Posts     []Post  `json:"posts"`
			}

			var forum Forum
			if err := cfg.DB.Preload("Server").Preload(
				"Posts", utils.Paginate(c, nil), s.Preload("User"),
			).Take(&forum, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("不存在此论坛", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("获取论坛数据失败", nil)
			} else {
				return 200, Res("", &forum)
			}
		},
	)
}
