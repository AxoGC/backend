package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuyGoods(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/goods/:id/buy", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.URI | p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID    uint `uri:"id"`
			Times uint `json:"times"`
		}) (int, *utils.Resp) {

			var goods []utils.Good
			if err := cfg.DB.Take(&prop, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("不存在该商品", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找商品失败", nil)
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

			return 200, Res("商品购买成功", nil)
		},
	)
}
