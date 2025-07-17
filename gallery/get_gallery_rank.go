package main

import "github.com/gin-gonic/gin"

func GetGalleryRank(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/gallery/rank", func(c *gin.Context) {

		type Album struct {
			ID    uint   `json:"id"`
			Label string `json:"label"`
			Likes uint   `json:"likes"`
		}

		var albums []Album
		if err := cfg.DB.Table("albums").
			Select("albums.id as id, albums.label as label, COALESCE(SUM(images.likes), 0) as likes").
			Joins("LEFT JOIN images ON images.album_id = albums.id").
			Group("albums.id").Find(&albums).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("统计相册点赞数据失败", nil))
			return
		}

		c.JSON(200, Res("", albums))
	}
}
