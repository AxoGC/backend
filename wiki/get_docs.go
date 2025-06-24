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

func GetDocs(cfg *p.Config) (string, string, []gin.HandlerFunc) {
	return "GET", "/docs/:slug", []gin.HandlerFunc{
		p.Preload(
			cfg, &p.Option{Bind: p.Uri}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug" binding:"required"`
			}) (int, *utils.Resp) {

				var doc struct {
					ID         uint      `json:"id"`
					CreatedAt  time.Time `json:"createdAt"`
					UpdatedAt  time.Time `json:"updatedAt"`
					Title      string    `json:"title"`
					Content    string    `json:"content"`
					DocGroupID uint      `json:"docGroupId"`
					DocGroup   struct {
						ID    uint   `json:"id"`
						Label string `json:"label"`
					} `json:"docGroup"`
					UserID uint `json:"userId"`
					User   struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
					Reviews []struct {
						ID             uint      `json:"id"`
						UpdatedAt      time.Time `json:"updatedAt"`
						Content        string    `json:"content"`
						ReviewableID   uint      `json:"reviewableId"`
						ReviewableType string    `json:"reviewableType"`
						UserID         uint      `json:"userId"`
						User           struct {
							ID   uint   `json:"id"`
							Name string `json:"name"`
						} `json:"user"`
					} `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:docs"`
					ReviewCount uint `json:"reviewCount"`
				}

				if err := cfg.DB.
					Model(new(utils.Doc)).
					Preload("DocGroup",
						s.Model(new(utils.DocGroup)),
					).
					Preload("User",
						s.Model(new(utils.User)),
					).
					Preload("Reviews",
						s.Model(new(utils.Review)),
						utils.Paginate(c, nil),
						s.Preload("User",
							s.Model(new(utils.User)),
						),
					).
					Take(&doc, "slug = ?", r.Slug).
					Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 400, Res("不存在此文档", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查询失败", nil)
				} else {
					return 200, Res("", &doc)
				}
			},
		),
	}
}
