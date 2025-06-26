package main

import (
	"github.com/axogc/backend/utils"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
)

func ListForumGroups(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/forum-groups", func(c *gin.Context) {

		var forumGroups []struct {
			ID     uint   `json:"id"`
			Label  string `json:"label"`
			Sort   int    `json:"sort"`
			Forums []struct {
				ID           uint   `json:"id"`
				ForumGroupID uint   `json:"forumGroupId"`
				Slug         string `json:"slug"`
				Title        string `json:"title"`
				SubTitle     string `json:"subTitle"`
				PostCount    uint   `json:"postCount"`
				Sort         int    `json:"sort"`
			}
		}
		if err := cfg.DB.Model(new(utils.ForumGroup)).Preload("Forums",
			s.Model(new(utils.Forum)),
		).Find(&forumGroups).Error; err != nil {
			c.JSON(500, Res("获取论坛列表失败", nil))
			c.Error(err)
		} else {
			c.JSON(200, Res("", forumGroups))
		}
	}
}
