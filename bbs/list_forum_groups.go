package main

import (
	"github.com/gin-gonic/gin"
)

func ListForumGroups(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/forum-groups", func(c *gin.Context) {

		type Forum struct {
			ID           uint   `json:"id"`
			ForumGroupID uint   `json:"forumGroupId"`
			Slug         string `json:"slug"`
			Title        string `json:"title"`
			SubTitle     string `json:"subTitle"`
			PostCount    uint   `json:"postCount"`
			Sort         int    `json:"sort"`
		}

		type ForumGroup struct {
			ID     uint    `json:"id"`
			Label  string  `json:"label"`
			Sort   int     `json:"sort"`
			Forums []Forum `json:"forums"`
		}

		var forumGroups []ForumGroup

		if err := cfg.DB.Preload("Forums").Find(&forumGroups).Error; err != nil {
			c.JSON(500, Res("获取论坛列表失败", nil))
			c.Error(err)
		} else {
			c.JSON(200, Res("", forumGroups))
		}
	}
}
