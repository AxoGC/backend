package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelReviews(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "DELETE", "/reviews/:id", p.Preload(
		cfg, &p.Option{Permission: p.Login, Bind: p.Uri}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID uint `uri:"id" binding:"required"`
		}) (int, error, *Resp) {

			var review utils.Review
			if err := cfg.DB.Take(&review, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, nil, &Resp{"找不到对应的评论", nil}
			} else if err != nil {
				return 500, err, &Resp{"查找评论失败", nil}
			} else if review.UserID != u.ID && !u.Admin {
				return 403, nil, &Resp{"你不是这个评论的发送者", nil}
			}

			if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

				if err := tx.Delete(&utils.Review{ID: r.ID}).Error; err != nil {
					return utils.NewTxData(500, err, "评论删除失败")
				}

				if err := tx.Table(review.ReviewableType).Where(review.ReviewableID).Update(
					"review_count", gorm.Expr("review_count - ?", 1),
				).Error; err != nil {
					return utils.NewTxData(500, err, "评论数量修改失败")
				}

				return nil
			}).(*utils.TxData); ok {
				return err.Code, err.Err, &Resp{err.Message, nil}
			}

			return 200, nil, &Resp{"评论删除成功", nil}
		},
	)
}
