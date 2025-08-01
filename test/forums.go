package main

import (
	"github.com/axogc/backend/utils"
	"github.com/samber/lo"
)

var TestForums = []utils.Forum{
	{
		ID:           1,
		Slug:         "chat",
		Title:        "聊天室", 
		SubTitle:     "一起来唠嗑∠( ᐛ 」∠)＿",
		Profile:      "这里是大家自由聊天的地方，无论是日常闲聊、分享心情，还是讨论热门话题，都欢迎大家参与。让我们在这里建立友谊，分享快乐！",
		ForumGroupID: 1,
		PostCount:    4,
	},
	{
		ID:           2,
		Slug:         "minecraft-bedrock",
		Title:        "我的世界基岩版", 
		SubTitle:     "为什么基岩版叫基岩版？",
		Profile:      "专门讨论我的世界基岩版（Minecraft Bedrock Edition）的版块。在这里可以分享建筑作品、交流游戏技巧、寻找联机伙伴，一起探索无限可能的方块世界！",
		ForumGroupID: 2,
		ServerID:     lo.ToPtr(uint(2)),
		PostCount:    2,
	},
	{
		ID:           3,
		Slug:         "minecraft-java",
		Title:        "我的世界Java版", 
		SubTitle:     "要致富，先撸树",
		Profile:      "我的世界Java版玩家的聚集地！讨论MOD推荐、红石科技、建筑设计、服务器推荐等。无论你是萌新还是老MC，这里都有你需要的攻略和伙伴。",
		ForumGroupID: 2,
		ServerID:     lo.ToPtr(uint(3)),
		PostCount:    4,
	},
	{
		ID:           4,
		Slug:         "dont-starve",
		Title:        "饥荒联机版", 
		SubTitle:     "浆果 好赤",
		Profile:      "饥荒联机版（Don't Starve Together）交流区域。分享生存技巧、角色攻略、BOSS打法，一起在这个充满挑战的世界中互助求生！",
		ServerID:     lo.ToPtr(uint(4)),
		ForumGroupID: 2,
		PostCount:    1,
	},
	{
		ID:           5,
		Slug:         "terraria",
		Title:        "泰拉瑞亚", 
		SubTitle:     "去吧，泰拉瑞亚战士！",
		Profile:      "泰拉瑞亚玩家的冒险基地！讨论装备合成、BOSS攻略、建筑设计、MOD推荐。在这个2D沙盒世界中，每个人都是勇敢的探险家！",
		ForumGroupID: 2,
		PostCount:    2,
		ServerID:     lo.ToPtr(uint(5)),
	},
	{
		ID:           6,
		Slug:         "stardew-valley",
		Title:        "星露谷物语", 
		SubTitle:     "小时候玩这个被紫色飞天苦茶追了",
		Profile:      "星露谷物语温馨交流区。分享农场设计、种植攻略、村民好感度提升、节日活动等。在这个宁静的乡村里，享受慢节奏的田园生活吧！",
		ForumGroupID: 2,
		PostCount:    1,
	},
	{
		ID:           7,
		Slug:         "sky",
		Title:        "光·遇",
		SubTitle:     "因光而遇的旅行…", 
		Profile:      "光·遇（Sky: Children of the Light）玩家交流天地。分享每日任务攻略、先祖位置、服装收集心得，一起在这个充满温暖的云端世界中传递光明！",
		ForumGroupID: 2,
		PostCount:    1,
	},
	{
		ID:           8,
		Slug:         "feedbacks",
		Title:        "BUG或意见反馈", 
		SubTitle:     "帮助我们变得更好 (๑•̀ㅂ•́)و✧",
		Profile:      "发现网站BUG或有改进建议？请在这里告诉我们！您的每一个反馈都是我们前进的动力，让我们一起打造更好的社区环境。",
		ForumGroupID: 3,
		PostCount:    2,
	},
	{
		ID:           9,
		Slug:         "reports",
		Title:        "玩家举报", 
		SubTitle:     "维护社区秩序，人人有责",
		Profile:      "如果您遇到违规行为、恶意用户或不当内容，请在此举报。我们会认真处理每一个举报，共同维护一个健康、友好的社区环境。",
		ForumGroupID: 3,
		PostCount:    2,
	},
}
