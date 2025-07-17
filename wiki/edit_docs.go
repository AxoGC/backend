package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditDocs(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "PUT", "/docs/:slug", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				OldSlug      string `uri:"slug"`
				Slug         string `json:"slug"`
				Title        string `json:"title"`
				Content      string `json:"content"`
				DocGroupSlug uint   `json:"docGroupSlug"`
				Sort         int    `json:"sort"`
			}) (int, *utils.Resp) {

				var docGroup utils.DocGroup
				if err := cfg.DB.Take(&docGroup, "slug = ?", r.DocGroupSlug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("找不到对应的文档组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找文档组失败", nil)
				}

				if result := cfg.DB.Model(new(utils.Doc)).Where("slug = ?", r.OldSlug).Updates(map[string]any{
					"slug":         r.Slug,
					"title":        r.Title,
					"content":      r.Content,
					"doc_group_id": docGroup.ID,
					"sort":         r.Sort,
					"user_id":      u.ID,
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
