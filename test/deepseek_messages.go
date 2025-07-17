package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestDeepSeekMessages = []utils.DeepSeekMessage{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, 0, -1),
		UserID:    1, // 史蒂夫
		Role:      "user",
		Content:   "请帮我设计一个自动化农场的红石电路",
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, 0, -1),
		UserID:    1,
		Role:      "assistant",
		Content:   "我可以帮您设计一个高效的自动化农场。首先，我们需要考虑作物类型...",
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, 0, -2),
		UserID:    1, // 艾莉克斯
		Role:      "user",
		Content:   "泰拉瑞亚的机械boss应该怎么打？",
	},
	{
		ID:        4,
		CreatedAt: time.Now().AddDate(0, 0, -2),
		UserID:    1,
		Role:      "assistant",
		Content:   "机械boss是困难模式的重要挑战。建议先准备好钛金装备...",
	},
}
