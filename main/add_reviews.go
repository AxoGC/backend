package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddReviews(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "POST", "/:type/:id/reviews", p.Preload(
		cfg, &p.Option{Permission: p.Login, Bind: p.Uri | p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Type     string `uri:"type" binding:"required,oneof=docs posts albums reviews"`
			ID       uint   `uri:"id" binding:"required"`
			Content  string `json:"content"`
			Attitude *bool  `json:"attitude"`
		}) (int, error, *Resp) {

			if err := cfg.DB.Table(r.Type).Select("1").Take(new(int), r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, nil, &Resp{"找不到对应的文档", nil}
			} else if err != nil {
				return 500, err, &Resp{"查找文档失败", nil}
			}

			if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

				if err := tx.Create(&utils.Review{
					Content:        r.Content,
					Attitude:       r.Attitude,
					UserID:         u.ID,
					ReviewableID:   r.ID,
					ReviewableType: r.Type,
				}).Error; err != nil {
					return utils.NewTxData(500, err, "创建评论失败")
				}

				if err := tx.Table(r.Type).Where(r.ID).Update("review_count", gorm.Expr("review_count + ?", 1)).Error; err != nil {
					return utils.NewTxData(500, err, "修改评论数量失败")
				}

				return nil
			}).(*utils.TxData); ok {
				return err.Code, err.Err, &Resp{err.Message, nil}
			}

			return 200, nil, &Resp{"评论发送成功", nil}
		},
	)
}
