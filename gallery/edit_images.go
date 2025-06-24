package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func EditImages(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/images/:id", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.Uri | p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				ID    uint   `uri:"id" binding:"required"`
				Label string `json:"label"`
			}) (int, *Resp) {

				// 需要权限检查器

				if result := cfg.DB.Where(&utils.Image{ID: r.ID}).Update("label", r.Label); result.RowsAffected == 0 {
					c.JSON(404, Resp{"找不到对应的图片", nil})
					return
				} else if result.Error != nil {
					c.JSON(500, Resp{"更新图片标题失败", nil})
					return
				}

				c.JSON(200, Resp{"更新成功", nil})
			},
		),
	}
}
