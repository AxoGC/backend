package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestPosts = []utils.Post{
	{
		ID:          1,
		CreatedAt:   time.Now().AddDate(0, 0, -5),
		UpdatedAt:   time.Now().AddDate(0, 0, -3),
		Pinned:      true,
		Title:       "欢迎新玩家加入我们的社区！",
		Slug:        "welcome-new-players",
		ForumID:     1, // 社区公告
		Content:     "# 欢迎来到游戏社区\n\n各位新老玩家大家好！欢迎加入我们这个温馨的游戏社区...",
		Markdown:    true,
		UserID:      1, // 管理员
		ReviewCount: 12,
	},
	{
		ID:          2,
		CreatedAt:   time.Now().AddDate(0, 0, -10),
		UpdatedAt:   time.Now().AddDate(0, 0, -8),
		Pinned:      false,
		Title:       "分享一下我的大型城堡建筑",
		Slug:        "my-castle-build-showcase",
		ForumID:     3, // 建筑作品展示
		Content:     "花了一个多月时间，终于完成了这座中世纪风格的大型城堡！\n\n包含完整的内部装饰和功能区域...",
		Markdown:    true,
		UserID:      1, // 史蒂夫
		ReviewCount: 28,
	},
	{
		ID:          3,
		CreatedAt:   time.Now().AddDate(0, 0, -7),
		UpdatedAt:   time.Now().AddDate(0, 0, -6),
		Pinned:      false,
		Title:       "红石全自动农场设计图分享",
		Slug:        "redstone-auto-farm-design",
		ForumID:     3, // 红石技术交流
		Content:     "## 全自动农场制作教程\n\n这个设计可以实现完全自动化的作物收割...",
		Markdown:    true,
		UserID:      1, // 史蒂夫
		ReviewCount: 15,
	},
	{
		ID:          4,
		CreatedAt:   time.Now().AddDate(0, 0, -12),
		UpdatedAt:   time.Now().AddDate(0, 0, -11),
		Pinned:      false,
		Title:       "泰拉瑞亚机械Boss攻略详解",
		Slug:        "terraria-mechanical-boss-guide",
		ForumID:     5, // 泰拉瑞亚讨论区
		Content:     "困难模式三大机械Boss的详细攻略来了！\n\n## 机械眼\n准备装备：钛金套装...",
		Markdown:    true,
		UserID:      1, // 艾莉克斯
		ReviewCount: 22,
	},
}
