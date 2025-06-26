package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditDocs(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "PATCH", "/docs/:slug", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"Roles", "Roles.Role"}}, nil,
		utils.WithRolesAuth(
			[]utils.Role{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				Uri        string `uri:"slug" binding:"required"`
				Slug       string `json:"slug" binding:"required"`
				Title      string `json:"title" binding:"required"`
				Content    string `json:"content" binding:"required"`
				DocGroupID uint   `json:"docGroupId" binding:"required"`
				Sort       int    `json:"sort"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Where("slug = ?", r.Uri).Updates(&utils.Doc{
					Slug:       r.Slug,
					Title:      r.Title,
					Content:    r.Content,
					DocGroupID: r.DocGroupID,
					Sort:       r.Sort,
					UserID:     u.ID,
				}); errors.Is(result.Error, gorm.ErrDuplicatedKey) {
					return 400, Res("已存在相同名称或标识的文档", nil)
				} else if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
					return 404, Res("不存在对应的文档组", nil)
				} else if result.RowsAffected == 0 {
					return 400, Res("找不到对应的文档", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 400, Res("更新文档失败", nil)
				} else {
					return 200, Res("更新文档组成功", nil)
				}
			},
		),
	)
}
