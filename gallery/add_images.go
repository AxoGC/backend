package main

import (
	"errors"
	"path/filepath"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddImages(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/album/:slug/images", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.Other}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Slug  string `uri:"slug" binding:"required"`
				Label string `form:"label" binding:"required,min=3,max=32"`
			}) (int, *utils.Resp) {

				var album utils.Album
				if err := cfg.DB.Take(&album, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("找不到对应的相册", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找相册失败", nil)
				}

				if album.Protected && !u.Admin && album.UserID != u.ID && (album.GuildID == nil || u.GuildID == nil || *album.GuildID != *u.GuildID) {
					return 403, Res("权限不足", nil)
				}

				fileHeader, err := c.FormFile("file")
				if err != nil {
					return 400, Res("无法读取图片文件", nil)
				}

				ext, err := utils.GetExtByFileHeader(fileHeader)
				if errors.Is(err, utils.ErrNotSupportedImageType) {
					return 400, Res("不支持的文件类型", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("解析图片文件失败", nil)
				}

				filename := uuid.New().String() + ext

				err = cfg.DB.Transaction(func(tx *gorm.DB) error {

					if err := c.SaveUploadedFile(
						fileHeader, filepath.Join(cfg.Env.FilePath, cfg.Env.ImagePath, filename),
					); err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("图片保存失败", nil))
					}

					if err := tx.Where(&album).Update("image_count", gorm.Expr("image_count + ?", 1)).Error; err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("相册图片数量修改失败", nil))
					}

					if err := tx.Create(&utils.Image{
						Filename: filename,
						AlbumID:  album.ID,
						Label:    r.Label,
						UserID:   u.ID,
					}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
						return utils.TxRes(409, Res("文件名冲突", nil))
					} else if err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("文件保存失败", nil))
					}

					return nil
				})

				if res, ok := err.(*utils.TxResp[utils.Resp]); ok {
					return res.Code, res.Data
				} else if err != nil {
					c.Error(err)
					return 500, Res("事务执行异常", nil)
				}

				return 200, Res("图片保存成功", nil)
			},
		),
	}
}
