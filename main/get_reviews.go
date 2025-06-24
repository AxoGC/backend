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

func GetReviews(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "GET", "/reviews/:id", p.Preload(
		cfg, &p.Option{Bind: p.Uri}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID uint `uri:"id" binding:"required"`
		}) (int, error, *Resp) {

			var review struct {
				ID        uint      `json:"id"`
				UpdatedAt time.Time `json:"updatedAt"`
				Content   string    `json:"content"`
				Attitude  *bool     `json:"attitude"`
				UserID    uint      `json:"userId"`
				User      struct {
					ID   uint   `json:"id"`
					Name string `json:"name"`
				} `json:"user"`
				ReviewCount uint `json:"reviewCount"`
				Reviews     []struct {
					ID        uint      `json:"id"`
					UpdatedAt time.Time `json:"updatedAt"`
					Content   string    `json:"content"`
					Attitude  *bool     `json:"attitude"`
					UserID    uint      `json:"userId"`
					User      struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
					ReviewCount    uint   `json:"reviewCount"`
					ReviewableID   uint   `json:"reviewableId"`
					ReviewableType string `json:"reviewableType"`
				} `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:reviews"`
			}
			if err := cfg.DB.Model(new(utils.Review)).Preload("User",
				s.Model(new(utils.User)),
			).Preload("Reviews",
				s.Model(new(utils.Review)),
				utils.Paginate(c, nil),
				s.Preload("User",
					s.Model(new(utils.User)),
				),
			).Take(&review, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, nil, &Resp{"找不到对应的评论", nil}
			} else if err != nil {
				return 500, err, &Resp{"查找评论失败", nil}
			}

			return 200, nil, &Resp{"", &review}
		},
	)
}
