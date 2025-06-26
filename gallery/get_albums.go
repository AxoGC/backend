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
	return "GET", "/albums", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Bind: p.Uri | p.Query, Permission: p.Auto}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				Slug string `uri:"slug" binding:"required"`
				All  bool   `form:"all"`
			}) (int, *utils.Resp) {

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
						ID             uint   `json:"id"`
						Content        string `json:"content"`
						UserID         uint   `json:"userId"`
						ReviewableID   uint   `json:"reviewableId"`
						ReviewableType string `json:"reviewableType"`
						User           struct {
							ID   uint   `json:"id"`
							Name string `json:"name"`
						}
					} `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:albums"`
				}

				query := cfg.DB.Model(new(utils.Album)).Preload("User").Preload("Images",
					utils.Paginate(c, &utils.PaginateConfig{KeyPrefix: "images"}),
				).Preload("Reviews",
					utils.Paginate(c, &utils.PaginateConfig{KeyPrefix: "users"}),
					s.Preload("User"),
				)

				if !r.All || !(u != nil && u.Admin) {
					query = query.Where("private = ?", false)
				}

				if err := query.First(&album, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 400, Res("不存在这个相册", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("获取相册失败", nil)
				}

				return 200, Res("", &album)
			},
		),
	}
}
