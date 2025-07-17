package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func ListAlbums(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/albums", p.Preload(
		cfg.Config, &p.Option{Bind: p.Query}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
		}) (int, *utils.Resp) {

			type Album struct {
				ID         uint   `json:"id"`
				Label      string `json:"label"`
				Slug       string `json:"slug"`
				ImageCount uint   `json:"imageCount"`
			}

			var albums []Album
			if err := cfg.DB.
				Model(new(utils.Album)).
				Scopes(utils.Paginate(c, nil)).
				Find(&albums).
				Error; err != nil {
				c.Error(err)
				return 500, Res("获取相册列表失败", nil)
			}

			return 200, Res("", albums)
		},
	)
}
