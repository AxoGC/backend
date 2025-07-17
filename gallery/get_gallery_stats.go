package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetGalleryStats(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/gallery/stats", func(c *gin.Context) {

		var albumCount int64
		if err := cfg.DB.Model(new(utils.Album)).Count(&albumCount).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("统计相册数量失败", nil))
			return
		}

		var imageCount int64
		if err := cfg.DB.Model(new(utils.Image)).Count(&imageCount).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("统计图片数量失败", nil))
			return
		}

		var reviewCount int64
		if err := cfg.DB.Model(new(utils.Review)).Where("reviewable_type = ?", "albums").Count(&reviewCount).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("统计评论数量失败", nil))
			return
		}

		var likesCount int64
		if err := cfg.DB.Model(new(utils.Image)).Pluck("SUM(likes) as likes_count", &likesCount).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("统计点赞数量失败", nil))
			return
		}

		c.JSON(200, Res("", gin.H{
			"albumCount":  albumCount,
			"imageCount":  imageCount,
			"reviewCount": reviewCount,
			"likesCount":  likesCount,
		}))
	}
}
