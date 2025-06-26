package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/albums/:slug", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug"`
			}) (int, *utils.Resp) {

				var album utils.Album
				if err := cfg.DB.Take(&album, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("没有对应的相册", nil)
				} else if err != nil {
					return 500, Res("查找相册失败", nil) 
				} else if !u.Admin && u.ID != album.UserID && (u.GuildID == nil || album.GuildID == nil || *u.GuildID != *album.GuildID) {
					return 403, Res("你没有权限删除相册", nil)
				}

				return 200, Res("相册删除成功", nil)
			},
		),
	}
}
