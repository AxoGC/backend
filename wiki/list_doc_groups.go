package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
)

func ListDocGroups(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "GET", "/doc-groups", func(c *gin.Context) {

		type Doc struct {
			ID         uint   `json:"id"`
			Title      string `json:"title"`
			DocGroupID uint   `json:"docGroupId"`
			Sort       int    `json:"sort"`
			Slug       string `json:"slug"`
		}

		type DocGroup struct {
			ID    uint   `json:"id"`
			Label string `json:"label"`
			Slug  string `json:"slug"`
			Sort  int    `json:"sort"`
			Docs  []Doc  `json:"docs"`
		}

		var docGroups []DocGroup

		if err := cfg.DB.Preload("Docs",
			s.Order("sort desc"),
		).Find(&docGroups).Error; err != nil {
			c.Error(err)
			c.JSON(500, Res("获取文档列表失败", nil))
		} else {
			c.JSON(200, Res("", docGroups))
		}
	}
}
