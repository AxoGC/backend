package main

import (
	"github.com/gin-gonic/gin"
)

func GetGalleryCarousel(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/gallery/carousel", func(c *gin.Context) {

		type User struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		type Image struct {
			ID       uint   `json:"id"`
			Filename string `json:"filename"`
			Label    string `json:"label"`
			Likes    uint   `json:"likes"`
			UserID   uint   `json:"userId"`
			User     User   `json:"user"`
		}

		var images []Image
		if err := cfg.DB.Preload("User").Order("likes desc").Limit(5).Find(&images).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("获取轮播图失败", nil))
			return
		}

		c.JSON(200, Res("", images))
	}
}
