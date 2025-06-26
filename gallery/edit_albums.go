package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func EditAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "PATCH", "/albums/:slug", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.Uri | p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				OldSlug      string  `uri:"slug" binding:"required"`
				IsGuildAlbum bool    `json:"isGuildAlbum"`
				GuildID      *uint   `json:"-"`
				Slug         *string `json:"slug"`
				Label        *string `json:"label"`
				Profile      *string `json:"profile"`
				Private      *bool   `json:"private"`
				Protected    *bool   `json:"protected"`
			}) (int, *utils.Resp) {

				var album utils.Album
				if err := cfg.DB.Take(&album, "slug = ?", r.OldSlug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("找不到对应的相册", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查找相册失败", nil)
				} else if !u.Admin && album.UserID != u.ID {
					return 403, Res("你没有权限修改该相册", nil)
				}

				album.GuildID = lo.Ternary(r.IsGuildAlbum, u.GuildID, nil)

				if err := cfg.DB.Model(new(utils.Album)).Updates(r).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("此标识已被其他相册使用", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("更新相册失败", nil)
				}

				return 200, Res("相册编辑成功", nil)
			},
		),
	}
}
