package main

import (
	"github.com/axogc/backend/utils"
	"github.com/samber/lo"
)

var TestDocGroups = []utils.DocGroup{
	{
		ID:    1,
		Label: "关于 Axolotland",
		Slug:  "about",
	},
	{
		ID:       2,
		Label:    "我的世界互通版",
		Slug:     "minecraft-paper",
		ServerID: lo.ToPtr(uint(1)),
	},
	{
		ID:       3,
		Label:    "我的世界基岩版",
		Slug:     "minecraft-bedrock",
		ServerID: lo.ToPtr(uint(2)),
	},
	{
		ID:       4,
		Label:    "我的世界模组版",
		Slug:     "minecraft-fabric",
		ServerID: lo.ToPtr(uint(3)),
	},
	{
		ID:       5,
		Label:    "饥荒联机版",
		Slug:     "dont-starve",
		ServerID: lo.ToPtr(uint(4)),
	},
	{
		ID:       6,
		Label:    "泰拉瑞亚",
		Slug:     "terraria",
		ServerID: lo.ToPtr(uint(5)),
	},
	{
		ID:    7,
		Label: "星露谷物语",
		Slug:  "stardew-valley",
	},
}
