package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func ListAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/albums", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Bind: p.Query}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
			}) (int, *Resp) {

				var albums []struct {
					User struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
				}
				if err := cfg.DB.Preload("User").Scopes(utils.Paginate(c, nil)).Find(&albums).Error; err != nil {
					c.JSON(500, Resp{"获取相册列表失败", nil})
					c.Error(err)
					return
				}

				c.JSON(200, Resp{"", albums})
			},
		),
	}
}
