package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func GetCarousels(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "GET", "/carousels", p.Preload(
		cfg, nil, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, error, *Resp) {

			var carousels []struct {
				ID       uint   `json:"id"`
				Filename string `json:"filename"`
				Label    string `json:"label"`
			}

			if err := cfg.DB.Model(new(utils.Image)).Joins("Album").Find(
				&carousels, "albums.path = ?", "carousel",
			).Error; err != nil {
				return 500, err, &Resp{"获取轮播图相册失败", nil}
			}

			return 200, nil, &Resp{"", carousels}
		},
	)
}
