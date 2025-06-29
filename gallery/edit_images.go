package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditImages(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "PATCH", "/images/:id", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.URI | p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			ID    uint   `uri:"id" binding:"required"`
			Label string `json:"label"`
		}) (int, *utils.Resp) {

			var image utils.Image
			if err := cfg.DB.Preload("Album").Take(&image, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的图片", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找图片失败", nil)
			} else if !u.HasAnyRole(utils.Admin, utils.GalleryAdmin) &&
				u.ID != image.UserID &&
				u.ID != image.Album.UserID {
				return 403, Res("你没有权限更新该图片", nil)
			}

			if err := cfg.DB.Where(&utils.Image{ID: r.ID}).Update("label", r.Label).Error; err != nil {
				c.Error(err)
				return 500, Res("更新图片标题失败", nil)
			}

			return 200, Res("更新成功", nil)
		},
	)
}
