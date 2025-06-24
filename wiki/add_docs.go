package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDocs(cfg *p.Config) (string, string, []gin.HandlerFunc) {
	return "POST", "/doc-groups/:docGroupId/docs", []gin.HandlerFunc{
		p.Preload(
			cfg, &p.Option{Permission: p.Admin, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Slug       string `json:"slug" binding:"required,min=3,alphanum"`
				Title      string `json:"title" binding:"required"`
				Content    string `json:"content" binding:"required"`
				DocGroupID uint   `uri:"docGroupId" binding:"required"`
				Sort       int    `json:"sort"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Create(&utils.Doc{
					Slug:       r.Slug,
					Title:      r.Title,
					Content:    r.Content,
					DocGroupID: r.DocGroupID,
					Sort:       r.Sort,
					UserID:     u.ID,
				}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("已存在相同名称或标识的文档", nil)
				} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return 404, Res("不存在对应的文档组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("创建文档失败", nil)
				} else {
					return 200, Res("文档创建成功", nil)
				}
			},
		),
	}
}
