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
		}) (int, *Resp) {

			var forum struct {
				ID        uint   `json:"id"`
				Title     string `json:"title"`
				SubTitle  string `json:"subTitle"`
				Profile   string `json:"profile"`
				PostCount uint   `json:"postCount"`
				Posts     []struct {
					ID        uint      `json:"id"`
					UpdatedAt time.Time `json:"updatedAt"`
					Pinned    bool      `json:"pinned"`
					Title     string    `json:"title"`
					ForumID   uint      `json:"forumId"`
					UserID    uint      `json:"userId"`
					User      struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
					ReviewCount uint `json:"reviewCount"`
				}
			}
			if err := cfg.DB.Model(new(utils.Forum)).Preload("Posts",
				s.Model(new(utils.Post)),
				utils.Paginate(c, nil),
				s.Preload("User",
					s.Model(new(utils.User)),
				),
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
