package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

type GetOnlineRequest struct {
	Period int      `form:"period"`
	Server []string `form:"server"`
}

func GetOnline(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/online", p.Preload(
		cfg.Config, &p.Option{Bind: p.Query}, &GetOnlineRequest{7, nil},
		func(c *gin.Context, u *utils.User, r *GetOnlineRequest) (int, *utils.Resp) {

			type Data struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Data []uint `json:"data"`
			}

			var datas []Data
			var servers []utils.Server
			query := cfg.DB
			if r.Server != nil {
				query = query.Where("slug IN ?", r.Server)
			}
			if err := query.Find(&servers).Error; err != nil {
				return 500, Res("获取服务器列表失败", nil)
			}

			for _, server := range servers {
				data := Data{
					Name: server.Label,
					Type: "line",
				}
				if err := cfg.DB.Model(new(utils.Online)).Where(
					"server_id = ?", server.ID,
				).Limit(r.Period).Pluck("count", &data.Data).Error; err != nil {
					c.Error(err)
					return 500, Res("获取在线数据失败", nil)
				}
				datas = append(datas, data)
			}

			return 200, Res("", datas)
		},
	)
}
