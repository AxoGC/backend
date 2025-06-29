package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func ListGoods(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "GET", "/goods", []gin.HandlerFunc{
		func(c *gin.Context) {

			var goods []utils.Good
			if err := cfg.DB.Preload("Prop").Scopes(utils.Paginate(c, nil)).Find(&goods).Error; err != nil {
				c.JSON(500, Res("获取商品列表失败", nil))
				c.Error(err)
				return
			}

			c.JSON(200, Res("", goods))
		},
	}
}
