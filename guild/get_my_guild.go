package main

/*

func GetMyGuild(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/my", p.Preload(
		cfg.Config, &p.Option{Login: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *utils.Resp) {

			var guild struct {
				CreatedAt time.Time `json:"createdAt"`
				Name      string    `json:"name"`
				Path      string    `json:"path"`
				Count     uint      `json:"count"`
				Profile   string    `json:"profile"`
				Notice    string    `json:"notice"`
				Money     uint      `json:"money"`
				Users     []struct {
					ID        uint   `json:"id"`
					Name      string `json:"name"`
					GuildID   *uint  `json:"guildId"`
					GuildRole uint   `json:"guildRole"`
				} `json:"users"`
			}

			if u.GuildID == nil {
				return 400, Res("你尚未加入任何公会", nil)
			}

			if err := cfg.DB.Where(&utils.Guild{ID: *u.GuildID}).Take(&guild).Error; err != nil {
				c.Error(err)
				return 500, Res("获取公会信息失败", nil)
			}

			return 200, Res("", &guild)
		},
	)
}
*/
