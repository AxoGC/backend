package main

import (
	"github.com/axogc/backend/utils"
	"github.com/samber/lo"
)

var TestUsers = []utils.User{
	{
		ID:        1,
		Name:      "Sauce233",
		Exp:       9999,
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMye",
		GenderID:  utils.Male,
		Profile:   "服主",
		DailyCoin: 1000,
		HonorCoin: 5000,
		Checkin:   365,
		Email:     "2892709432@qq.com",
		QQ:        lo.ToPtr("2892709432"),
		MCBEName:  lo.ToPtr("bestcb5843"),
		MCJEName:  lo.ToPtr("Bestcb233"),
		Setting:   map[string]any{},
	},
}
