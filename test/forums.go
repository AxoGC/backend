package main

import "github.com/axogc/backend/utils"

/*
var Forums = []utils.ForumGroup{
	{ID: 1, Label: "闲聊分区", Forums: []utils.Forum{
		{ID: 1, Slug: "chat", Title: "聊天室", SubTitle: "", Profile: "", Posts: []utils.Post{}},
	}},
	{ID: 2, Label: "游戏交流分区", Forums: []utils.Forum{
		{ID: 2, Slug: "minecraft-bedrock", Title: "我的世界基岩版", SubTitle: "", Profile: ""},
		{ID: 3, Slug: "minecraft-java", Title: "我的世界Java版", SubTitle: "", Profile: ""},
		{ID: 4, Slug: "dont-starve", Title: "饥荒联机版", SubTitle: "", Profile: ""},
		{ID: 5, Slug: "terraria", Title: "泰拉瑞亚", SubTitle: "", Profile: ""},
		{ID: 6, Slug: "stardew-valley", Title: "星露谷物语", SubTitle: "", Profile: ""},
		{ID: 7, Slug: "sky", Title: "", SubTitle: "光遇", Profile: ""},
	}},
	{ID: 3, Label: "反馈分区", Forums: []utils.Forum{
		{ID: 8, Slug: "feedbacks", Title: "BUG或意见反馈", SubTitle: "", Profile: ""},
		{ID: 9, Slug: "reports", Title: "玩家举报", SubTitle: "", Profile: ""},
	}},
}
*/

var TestForums = []utils.Forum{
	{
		ID:           1,
		ForumGroupID: 1, // 游戏讨论
		Slug:         "minecraft-general",
		Title:        "我的世界综合讨论",
		SubTitle:     "Java版&基岩版通用讨论区",
		Profile:      "讨论我的世界游戏相关的所有内容，包括游戏技巧、建筑展示、MOD推荐等",
		PostCount:    156,
		Sort:         1,
		ServerID:     uintPtr(1), // 关联生存服务器
	},
	{
		ID:           2,
		ForumGroupID: 1,
		Slug:         "terraria-discussion",
		Title:        "泰拉瑞亚讨论区",
		SubTitle:     "2D沙盒冒险游戏",
		Profile:      "泰拉瑞亚游戏攻略、装备制作、Boss战技巧等内容讨论",
		PostCount:    89,
		Sort:         2,
		ServerID:     uintPtr(4), // 关联泰拉瑞亚服务器
	},
	{
		ID:           3,
		ForumGroupID: 2, // 技术交流
		Slug:         "redstone-tech",
		Title:        "红石技术交流",
		SubTitle:     "电路设计与自动化装置",
		Profile:      "分享红石电路设计、自动化装置制作教程和技术讨论",
		PostCount:    67,
		Sort:         1,
		ServerID:     nil,
	},
	{
		ID:           4,
		ForumGroupID: 3, // 作品展示
		Slug:         "building-showcase",
		Title:        "建筑作品展示",
		SubTitle:     "展示你的创意建筑",
		Profile:      "展示各种游戏中的精美建筑作品，分享建筑经验和技巧",
		PostCount:    203,
		Sort:         1,
		ServerID:     nil,
	},
	{
		ID:           5,
		ForumGroupID: 4, // 社区事务
		Slug:         "community-announcements",
		Title:        "社区公告",
		SubTitle:     "重要通知与公告",
		Profile:      "发布社区重要通知、活动公告、规则更新等信息",
		PostCount:    25,
		Sort:         1,
		ServerID:     nil,
	},
	{
		ID:           6,
		ForumGroupID: 5, // 服务器相关
		Slug:         "server-support",
		Title:        "服务器技术支持",
		SubTitle:     "服务器问题求助",
		Profile:      "服务器连接问题、游戏BUG反馈、技术支持等",
		PostCount:    78,
		Sort:         1,
		ServerID:     nil,
	},
}
