package main

/*

func DelGuilds(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/disband", p.Preload(
		cfg.Config, &p.Option{Login: p.Login}, nil,
		WithGuildRolesAuth(
			[]utils.DictID{Owner},
			func(c *gin.Context, u *utils.User, r *struct{}) (int, *utils.Resp) {

				if err := cfg.DB.Delete(utils.Guild{ID: *u.GuildID}).Error; err != nil {
					c.Error(err)
					return 500, Res("公会解散失败", nil)
				}

				return 200, Res("公会解散成功", nil)
			},
		),
	)
}

*/
