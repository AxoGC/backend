package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditForums(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "PUT", "/forums/:slug", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.URI | p.JSON}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				OldSlug      string `uri:"slug"`
				Slug         string `json:"slug"`
				ForumGroupID uint   `json:"forumGroupId"`
				Title        string `json:"title"`
				SubTitle     string `json:"subTitle"`
				Profile      string `json:"profile"`
				Sort         int    `json:"sort"`
				ServerID     *uint  `json:"serverId"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Model(new(utils.Forum)).Where("slug = ?", r.OldSlug).Updates(map[string]any{
					"slug":           r.Slug,
					"forum_group_id": r.ForumGroupID,
					"title":          r.Title,
					"sub_title":      r.SubTitle,
					"profile":        r.Profile,
					"sort":           r.Sort,
					"server_id":      r.ServerID,
				}); errors.Is(result.Error, gorm.ErrCheckConstraintViolated) {
					return 404, Res("不存在对应的论坛组", nil)
				} else if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
					return 409, Res("该标题或标识已被其他论坛使用", nil)
				} else if result.RowsAffected == 0 {
					return 404, Res("不存在该论坛", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("论坛修改失败", nil)
				} else {
					return 200, Res("论坛修改成功", nil)
				}
			},
		),
	)
}
