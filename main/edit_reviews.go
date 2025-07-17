package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditReviews(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "PATCH", "/reviews/:id", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID       uint   `uri:"id" binding:"required"`
			Attitude *bool  `json:"attitude" binding:"required"`
			Content  string `json:"content"`
		}) (int, *utils.Resp) {

			var review utils.Review
			if err := cfg.DB.Take(&review, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的评论", nil)
			} else if err != nil {
				return 500, Res("查找评论失败", nil)
			} else if review.UserID != u.ID && !u.HasAnyRole(utils.Admin, utils.ReviewAdmin) {
				return 403, Res("你没有权限修改此评论", nil)
			}

			if err := cfg.DB.Where(&utils.Review{ID: r.ID}).Select("attitude", "content").Updates(r).Error; err != nil {
				return 500, Res("评论修改失败", nil)
			}

			return 200, Res("评论修改成功", nil)
		},
	)
}
