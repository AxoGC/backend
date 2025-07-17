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

func GetDocs(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "GET", "/docs/:slug", p.Preload(
		cfg, &p.Option{Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug" binding:"required"`
		}) (int, *utils.Resp) {

			type DocGroup struct {
				ID    uint   `json:"id"`
				Label string `json:"label"`
			}

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type Review struct {
				ID             uint      `json:"id"`
				UpdatedAt      time.Time `json:"updatedAt"`
				Content        string    `json:"content"`
				ReviewableID   uint      `json:"reviewableId"`
				ReviewableType string    `json:"reviewableType"`
				UserID         uint      `json:"userId"`
				User           User      `json:"user"`
			}

			type Doc struct {
				ID          uint      `json:"id"`
				CreatedAt   time.Time `json:"createdAt"`
				UpdatedAt   time.Time `json:"updatedAt"`
				Title       string    `json:"title"`
				Slug        string    `json:"slug"`
				Content     string    `json:"content"`
				Sort        int       `json:"sort"`
				DocGroupID  uint      `json:"docGroupId"`
				DocGroup    DocGroup  `json:"docGroup"`
				UserID      uint      `json:"userId"`
				User        User      `json:"user"`
				Reviews     []Review  `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:docs"`
				ReviewCount uint      `json:"reviewCount"`
			}

			var doc Doc

			if err := cfg.DB.Preload("DocGroup").Preload("User").Preload(
				"Reviews", utils.Paginate(c, nil), s.Preload("User"),
			).Take(&doc, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 400, Res("不存在此文档", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查询失败", nil)
			} else {
				return 200, Res("", &doc)
			}
		},
	)
}
