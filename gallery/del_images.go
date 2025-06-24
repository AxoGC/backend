package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func DelImages(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "DELETE", "/images", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				IDs []uint `json:"ids" binding:"required"`
			}) (int, *Resp) {

				if err := cfg.DB.Delete(new(utils.Image), r.IDs).Error; err != nil {
					c.JSON(500, Resp{"删除文件列表失败", nil})
					c.Error(err)
					return
				}

				c.JSON(200, Resp{"删除成功", nil})
			},
		),
	}
}
