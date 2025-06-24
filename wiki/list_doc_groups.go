package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
)

func ListDocGroups(cfg *p.Config) (string, string, []gin.HandlerFunc) {
	return "GET", "/doc-groups", []gin.HandlerFunc{
		func(c *gin.Context) {

			var docGroups []struct {
				ID    uint   `json:"id"`
				Label string `json:"label"`
				Sort  int    `json:"sort"`
				Docs  []struct {
					ID         uint   `json:"id"`
					Title      string `json:"title"`
					DocGroupID uint   `json:"docGroupId"`
					Sort       int    `json:"sort"`
					Slug       string `json:"slug"`
				} `json:"docs"`
			}

			if err := cfg.DB.Model(new(utils.DocGroup)).Preload("Docs",
				s.Model(new(utils.Doc)),
				s.Order("sort desc"),
			).Find(&docGroups).Error; err != nil {
				c.Error(err)
				c.JSON(500, Res("获取文档列表失败", nil))
			} else {
				c.JSON(200, Res("", docGroups))
			}
		},
	}
}
