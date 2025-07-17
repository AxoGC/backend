package main

import "github.com/axogc/backend/utils"

var TestForumGroups = []utils.ForumGroup{
	{
		ID:    1,
		Label: "闲聊分区",
		Slug:  "chat",
		Sort:  0,
	},
	{
		ID:    2,
		Label: "游戏交流分区",
		Slug:  "game",
		Sort:  0,
	},
	{
		ID:    3,
		Label: "反馈分区",
		Slug:  "feedback",
		Sort:  0,
	},
}
