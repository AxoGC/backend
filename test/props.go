package main

import "github.com/axogc/backend/utils"

var TestProps = []utils.Prop{
	{
		ID:          utils.GoldPouch,
		Label:       "金币袋",
		Description: "随机开出一定数额金币的袋子",
	},
	{
		ID:          utils.RaffleTicket,
		Label:       "抽奖券",
		Description: "可以在抽奖转盘使用",
	},
	{
		ID:          utils.ExpBoostCard,
		Label:       "经验加成卡",
		Description: "随机提升一定经验",
	},
	{
		ID:          utils.MysteryBox,
		Label:       "盲盒",
		Description: "包含随机道具的神秘盒子，开启后可获得意外惊喜",
	},
	{
		ID:          utils.PostPinCard,
		Label:       "帖子置顶卡",
		Description: "在一定时间段内置顶你的帖子",
	},
	{
		ID:          utils.GuildPinCard,
		Label:       "公会置顶卡",
		Description: "在一定时间段内置顶你的公会",
	},
	{
		ID:          utils.RenameCard,
		Label:       "改名卡",
		Description: "用于修改用户名",
	},
	{
		ID:          utils.RemakeCard,
		Label:       "重开卡",
		Description: "用于重开",
	},
}
