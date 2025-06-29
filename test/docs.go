package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

/*
import "github.com/axogc/backend/utils"

var Data = []utils.DocGroup{
	{ID: 1, Label: "关于Axolotland", Docs: []utils.Doc{
		{ID: 1, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
	{ID: 2, Label: "我的世界主服务器", Docs: []utils.Doc{
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
	{ID: 3, Label: "我的世界基岩版", Docs: []utils.Doc{
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
	{ID: 4, Label: "我的世界Java版（模组）", Docs: []utils.Doc{
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
	{ID: 5, Label: "饥荒联机版", Docs: []utils.Doc{
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
	{ID: 6, Label: "泰拉瑞亚", Docs: []utils.Doc{
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
	{ID: 7, Label: "星露谷物语", Docs: []utils.Doc{
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
		{ID: 2, Slug: "", Title: "", UserID: 1, Content: ""},
	}},
}
*/

var TestDocs = []utils.Doc{
	{
		ID:          1,
		CreatedAt:   time.Now().AddDate(0, -3, 0),
		UpdatedAt:   time.Now().AddDate(0, 0, -10),
		Slug:        "minecraft-beginner-guide",
		Title:       "我的世界新手指南",
		DocGroupID:  1, // 新手指南
		UserID:      2, // 史蒂夫
		Content:     "## 欢迎来到我的世界\n\n这是一个完全由方块组成的世界...",
		Sort:        1,
		ReviewCount: 15,
	},
	{
		ID:          2,
		CreatedAt:   time.Now().AddDate(0, -2, 0),
		UpdatedAt:   time.Now().AddDate(0, 0, -5),
		Slug:        "redstone-basics",
		Title:       "红石电路基础",
		DocGroupID:  3, // 技术文档
		UserID:      2, // 史蒂夫
		Content:     "# 红石电路入门\n\n红石是我的世界中的电路系统...",
		Sort:        1,
		ReviewCount: 8,
	},
	{
		ID:          3,
		CreatedAt:   time.Now().AddDate(0, -1, 0),
		UpdatedAt:   time.Now().AddDate(0, 0, -2),
		Slug:        "community-rules",
		Title:       "社区管理规则",
		DocGroupID:  4, // 社区规则
		UserID:      1, // 管理员
		Content:     "## 社区行为准则\n\n1. 尊重他人\n2. 禁止恶意行为...",
		Sort:        1,
		ReviewCount: 3,
	},
	{
		ID:          4,
		CreatedAt:   time.Now().AddDate(0, 0, -20),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
		Slug:        "terraria-boss-guide",
		Title:       "泰拉瑞亚Boss攻略",
		DocGroupID:  2, // 游戏攻略
		UserID:      3, // 艾莉克斯
		Content:     "# 泰拉瑞亚Boss战攻略\n\n## 克苏鲁之眼\n这是第一个boss...",
		Sort:        1,
		ReviewCount: 25,
	},
}
