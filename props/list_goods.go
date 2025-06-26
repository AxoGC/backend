package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func ListGoods(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "GET", "/goods", []gin.HandlerFunc{
		func(c *gin.Context) {

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
	}
}
