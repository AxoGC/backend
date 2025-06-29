package main

import "github.com/axogc/backend/utils"

var TestGames = []utils.Game{
	{
		ID:          utils.MinecraftBedrock,
		Label:       "我的世界基岩版",
		Description: "支持跨平台多人联机的沙盒游戏基岩版本",
	},
	{
		ID:          utils.MinecraftJava,
		Label:       "我的世界Java版",
		Description: "PC端经典版本，支持MOD和丰富的自定义内容",
	},
	{
		ID:          utils.DontStarve,
		Label:       "饥荒",
		Description: "充满挑战的生存冒险游戏",
	},
	{
		ID:          utils.Terraria,
		Label:       "泰拉瑞亚",
		Description: "2D沙盒动作冒险游戏",
	},
	{
		ID:          utils.StardewValley,
		Label:       "星露谷物语",
		Description: "温馨的农场模拟经营游戏",
	},
	{
		ID:          utils.Palworld,
		Label:       "幻兽帕鲁",
		Description: "开放世界生存制作游戏",
	},
	{
		ID:          utils.ARKSurvivalEvolved,
		Label:       "方舟：生存进化",
		Description: "恐龙主题的生存游戏",
	},
}
