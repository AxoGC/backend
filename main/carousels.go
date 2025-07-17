package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func GetCarousels(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "GET", "/carousels", p.Preload(
		cfg, nil, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *utils.Resp) {

			type Album struct {
				ID   uint   `json:"id"`
				Slug string `json:"slug"`
			}

			type Image struct {
				ID       uint   `json:"id"`
				Filename string `json:"filename"`
				Label    string `json:"label"`
				AlbumID  uint   `json:"albumId"`
				Album    Album  `json:"album"`
			}

			var carousels []Image

			if err := cfg.DB.Joins("Album").Find(
				&carousels, "Album.slug = ?", "carousels",
			).Error; err != nil {
				return 500, Res("获取轮播图相册失败", nil)
			}

			return 200, Res("", carousels)
		},
	)
}
