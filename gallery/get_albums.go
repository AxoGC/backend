package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/albums", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Bind: p.Uri | p.Query}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Path string `uri:"path" binding:"required"`
			}) (int, *Resp) {

				var album struct {
					User struct {
						ID   uint   `json:"id"`
						Name string `json:"name"`
					} `json:"user"`
					Images []struct {
						ID       uint   `json:"id"`
						Filename string `json:"filename"`
						Label    string `json:"label"`
						Likes    uint   `json:"likes"`
						AlbumID  uint   `json:"albumId"`
					} `json:"images"`
					Reviews []struct {
						ID        uint   `json:"id"`
						Content   string `json:"content"`
						UserID    uint   `json:"userId"`
						ReferID   uint   `json:"referId"`
						ReferType string `json:"referType"`
						User      struct {
							ID   uint   `json:"id"`
							Name string `json:"name"`
						}
					} `json:"reviews"`
				}

				if err := cfg.DB.Model(new(utils.Album)).Preload("User").Preload("Images",
					utils.Paginate(c, &utils.PaginateConfig{KeyPrefix: "images"}),
				).Preload("Reviews",
					utils.Paginate(c, &utils.PaginateConfig{KeyPrefix: "users"}),
					s.Preload("User"),
				).First(&album, "path = ?", r.Path).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(400, Resp{"不存在这个相册", nil})
					return
				} else if err != nil {
					c.JSON(500, Resp{"获取相册失败", nil})
					c.Error(err)
					return
				}

				c.JSON(200, Resp{"", &album})
			},
		),
	}
}
