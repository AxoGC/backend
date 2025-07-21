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
		GenderID:  nil,
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
	{
		ID:       2,
		Name:     "Nerakolo",
		GenderID: lo.ToPtr(utils.Femboy),
		Email:    "nerakolo@outlook.com",
	},
	{
		ID:    3,
		Name:  "Steve",
		Email: "steve@gmail.com",
	},
	{
		ID:    4,
		Name:  "Alex",
		Email: "alex@gmail.com",
	},
	{
		ID:      5,
		Name:    "Wilson",
		Email:   "wilson@gmail.com",
		Profile: "这是科学！",
	},
	{
		ID:      6,
		Name:    "Willow",
		Email:   "willow@gmail.com",
		Profile: "我能烧掉它吗？",
	},
	{
		ID:      7,
		Name:    "Wolfgang",
		Email:   "wolfgang@gmail.com",
		Profile: "它很弱，就像没吃饱的沃尔夫冈一样。",
	},
	{
		ID:      8,
		Name:    "Wendy",
		Email:   "wendy@gmail.com",
		Profile: "它看起来好悲伤。",
	},
}
