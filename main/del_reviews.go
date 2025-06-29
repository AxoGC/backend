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
		cfg, &p.Option{Login: p.Login, Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID uint `uri:"id" binding:"required"`
		}) (int, *utils.Resp) {

			var review utils.Review
			if err := cfg.DB.Take(&review, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的评论", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找评论失败", nil)
			} else if review.UserID != u.ID && !u.HasAnyRole(utils.RoleAdmin, utils.RoleReviewAdmin) {
				return 403, Res("你不是这个评论的发送者", nil)
			}

			if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

				if err := tx.Delete(&utils.Review{ID: r.ID}).Error; err != nil {
					return utils.TxRes(500, Res("评论删除失败", nil))
				}

				if err := tx.Table(review.ReviewableType).Where(review.ReviewableID).Update(
					"review_count", gorm.Expr("review_count - ?", 1),
				).Error; err != nil {
					return utils.TxRes(500, Res("评论数量修改失败", nil))
				}

				return nil
			}).(*utils.TxResp[utils.Resp]); ok {
				return err.Code, err.Data
			}

			return 200, Res("评论删除成功", nil)
		},
	)
}
