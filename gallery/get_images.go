package main

import (
	"errors"
	"time"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetImages(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "GET", "/images/:id", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Bind: p.Uri}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				ID uint `uri:"id" binding:"required"`
			}) (int, *utils.Resp) {

				var image struct {
					CreatedAt time.Time `json:"createdAt"`
					Filename  string    `json:"filename"`
					Label     string    `json:"label"`
					Likes     uint      `json:"likes"`
					UserID    uint      `json:"userId"`
					User      struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
					AlbumID uint `json:"albumId"`
					Album   struct {
						ID    uint   `json:"id"`
						Label string `json:"label"`
					} `json:"album"`
				}
				if err := cfg.DB.Model(new(utils.Image)).Preload("User",
					s.Model(new(utils.User)),
				).Preload("Album",
					s.Model(new(utils.Album)),
				).First(&image, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("不存在此图片", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("获取图片失败", nil)
				}

				return 200, Res("", &image)
			},
		),
	}
}
