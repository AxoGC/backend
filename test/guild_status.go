package main

import "github.com/axogc/backend/utils"

var TestGuildStatus = []utils.GuildStatus{
	{ID: utils.GuildBlocked, Label: "拉黑"},
	{ID: utils.GuildApplicant, Label: "申请中"},
	{ID: utils.GuildMember, Label: "公会成员"},
	{ID: utils.GuildAdmin, Label: "公会管理员"},
}
