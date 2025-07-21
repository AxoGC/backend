package main

/*

func Grant(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/grant", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON}, nil,
		WithGuildRolesAuth(
			[]utils.DictID{Owner},
			func(c *gin.Context, u *utils.User, r *struct {
				IDs []uint `json:"ids"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Where(
					"id IN ? AND guild_id IN ? AND guild_role = ?", r.IDs, u.GuildID, Member,
				).Update("guild_role", Admin).Error; err != nil {
					return 500, Res("提拔为管理员失败", nil)
				}

				return 200, Res("提拔为管理员成功", nil)
			},
		),
	)
}
*/
