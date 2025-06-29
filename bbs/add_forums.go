package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddForums(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/forum-groups/:forumGroupId/forums", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				ForumGroupID uint   `uri:"forumGroupId"`
				Slug         string `json:"slug"`
				Title        string `json:"title"`
				SubTitle     string `json:"subTitle"`
				Profile      string `json:"profile"`
				Sort         int    `json:"sort"`
				ServerID     *uint  `json:"serverId"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Model(new(utils.Forum)).Create(map[string]any{
					"forum_group_id": r.ForumGroupID,
					"slug":           r.Slug,
					"title":          r.Title,
					"sub_title":      r.SubTitle,
					"profile":        r.Profile,
					"sort":           r.Sort,
					"server_id":      r.ServerID,
				}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("已存在相同标识或标题的论坛", nil)
				} else if errors.Is(err, gorm.ErrCheckConstraintViolated) {
					return 422, Res("不存在对应的论坛组或服务器", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("创建论坛失败", nil)
				} else {
					return 200, Res("创建论坛成功", nil)
				}
			},
		),
	)
}
