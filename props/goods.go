package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterShops(r *gin.Engine, cfg *p.Config) {

	r.GET("/goods", p.Preload(
		cfg, &p.Option{}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) {

			var goods []struct {
				ID    uint   `json:"id"`
				Label string `json:"label"`
				Price uint   `json:"price"`
			}
			if err := cfg.DB.Model(new(utils.Prop)).Find(&goods, "price IS NOT NULL").Error; err != nil {
				c.JSON(500, Resp{"获取商品列表失败", nil})
				c.Error(err)
				return
			}

			c.JSON(200, Resp{"", goods})
		},
	))

	r.POST("/goods/:id/buy", p.Preload(
		cfg, &p.Option{Permission: p.Login, Bind: p.Uri | p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID    uint `uri:"id" binding:"required"`
			Count uint `json:"count" binding:"required"`
		}) {

			var prop utils.Prop
			if err := cfg.DB.Take(&prop, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(404, Resp{"不存在该商品", nil})
				return
			} else if err != nil {
				c.JSON(500, Resp{"查找商品失败", nil})
				c.Error(err)
				return
			} else if prop.Price == nil {
				c.JSON(400, Resp{"该商品尚不出售", nil})
				return
			}

			if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

				if err := u.CostCoin(tx, *prop.Price*r.Count); errors.Is(err, utils.ErrCoinNotEnough) {
					return utils.NewResErr(400, "金额不足")
				} else if err != nil {
					c.Error(err)
					return utils.NewResErr(500, "扣费失败")
				}

				return nil
			}).(*utils.ResErr); ok {
				c.JSON(err.Code, Resp{err.Message, nil})
				return
			}
		},
	))
}
