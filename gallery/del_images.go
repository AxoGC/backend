package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelImages(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "DELETE", "/images/:id", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				ID uint `json:"id" binding:"required"`
			}) (int, *utils.Resp) {

				var image utils.Image
				if err := cfg.DB.Preload("Album").Take(&image, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("找不到对应的图片", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找图片失败", nil)
				}

				if !u.Admin && u.ID != image.Album.UserID && u.ID != image.UserID {
					return 403, Res("你没有权限删除", nil)
				}

				err := cfg.DB.Transaction(func(tx *gorm.DB) error {

					if err := os.Remove(filepath.Join(cfg.Env.FilePath, cfg.Env.ImagePath, image.Filename)); err != nil && !errors.Is(err, os.ErrNotExist) {
						c.Error(err)
						return utils.TxRes(500, Res("删除图片文件失败", nil))
					}

					if err := tx.Delete(&utils.Image{ID: r.ID}).Error; err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("删除图片记录失败", nil))
					}

					if err := tx.Where(&image.Album).Update("image_count", gorm.Expr("image_count - ?", 1)).Error; err != nil {
						return utils.TxRes(500, Res("修改相册图片数量失败", nil))
					}

					return nil
				})

				if res, ok := err.(*utils.TxResp[utils.Resp]); ok {
					return res.Code, res.Data
				} else if err != nil {
					c.Error(err)
					return 500, Res("事务执行错误", nil)
				}

				return 200, Res("删除成功", nil)
			},
		),
	}
}
