package main

import (
	"time"

	"github.com/axogc/backend/utils"
	"github.com/samber/lo"
)

var TestReviews = []utils.Review{
	// 对帖子的评论 (posts)

	// 帖子1: 欢迎新玩家加入我们的社区! (聊天室)
	{ID: 1, UpdatedAt: time.Now().Add(1 * time.Hour), Content: "欢迎欢迎！希望大家都能在这里找到志同道合的朋友~", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 1, ReviewableType: "posts", ReviewCount: 0},
	{ID: 2, UpdatedAt: time.Now().Add(2 * time.Hour), Content: "新人报到！请多多指教！", Attitude: nil, UserID: 4, ReviewableID: 1, ReviewableType: "posts", ReviewCount: 0},
	{ID: 3, UpdatedAt: time.Now().Add(3 * time.Hour), Content: "社区氛围真的很棒，管理员们辛苦了", Attitude: lo.ToPtr(true), UserID: 5, ReviewableID: 1, ReviewableType: "posts", ReviewCount: 0},
	{ID: 4, UpdatedAt: time.Now().Add(4 * time.Hour), Content: "希望能多举办一些活动", Attitude: nil, UserID: 6, ReviewableID: 1, ReviewableType: "posts", ReviewCount: 0},

	// 帖子2: 分享一下我的大型城堡建筑 (我的世界Java版)
	{ID: 5, UpdatedAt: time.Now().Add(5 * time.Hour), Content: "哇！建筑技术太强了，能分享一下建筑思路吗？", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},
	{ID: 6, UpdatedAt: time.Now().Add(6 * time.Hour), Content: "这个城堡的细节做得很用心，特别是塔楼部分", Attitude: lo.ToPtr(true), UserID: 3, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},
	{ID: 7, UpdatedAt: time.Now().Add(7 * time.Hour), Content: "求教程！我也想学习这种建筑风格", Attitude: nil, UserID: 7, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},
	{ID: 8, UpdatedAt: time.Now().Add(8 * time.Hour), Content: "花了多长时间建造的？", Attitude: nil, UserID: 8, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},
	{ID: 9, UpdatedAt: time.Now().Add(9 * time.Hour), Content: "内饰装修也做了吗？", Attitude: nil, UserID: 2, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},

	// 帖子3: 红石全自动农场设计图分享 (我的世界Java版)
	{ID: 10, UpdatedAt: time.Now().Add(10 * time.Hour), Content: "这个红石电路太复杂了，有简化版本吗？", Attitude: nil, UserID: 4, ReviewableID: 3, ReviewableType: "posts", ReviewCount: 0},
	{ID: 11, UpdatedAt: time.Now().Add(11 * time.Hour), Content: "效率很高！已经在我的服务器上搭建了", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 3, ReviewableType: "posts", ReviewCount: 0},
	{ID: 12, UpdatedAt: time.Now().Add(12 * time.Hour), Content: "红石大佬！膜拜", Attitude: lo.ToPtr(true), UserID: 5, ReviewableID: 3, ReviewableType: "posts", ReviewCount: 0},
	{ID: 13, UpdatedAt: time.Now().Add(13 * time.Hour), Content: "能出个视频教程吗？图片看不太懂", Attitude: nil, UserID: 8, ReviewableID: 3, ReviewableType: "posts", ReviewCount: 0},

	// 帖子4: 泰拉瑞亚机械Boss攻略详解 (泰拉瑞亚)
	{ID: 14, UpdatedAt: time.Now().Add(14 * time.Hour), Content: "攻略很详细，按照这个打法成功通关了！", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 15, UpdatedAt: time.Now().Add(15 * time.Hour), Content: "机械骷髅王还是很难打，需要多练习", Attitude: nil, UserID: 3, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 16, UpdatedAt: time.Now().Add(16 * time.Hour), Content: "推荐的武器搭配很实用", Attitude: lo.ToPtr(true), UserID: 7, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 17, UpdatedAt: time.Now().Add(17 * time.Hour), Content: "场地布置的建议很棒！", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},

	// 帖子5: 克苏鲁之眼简单打法分享 (泰拉瑞亚)
	{ID: 18, UpdatedAt: time.Now().Add(18 * time.Hour), Content: "新手友好的攻略，感谢分享！", Attitude: lo.ToPtr(true), UserID: 4, ReviewableID: 5, ReviewableType: "posts", ReviewCount: 0},
	{ID: 19, UpdatedAt: time.Now().Add(19 * time.Hour), Content: "这确实是最简单的打法了", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 5, ReviewableType: "posts", ReviewCount: 0},
	{ID: 20, UpdatedAt: time.Now().Add(20 * time.Hour), Content: "木弓真的够用吗？", Attitude: nil, UserID: 8, ReviewableID: 5, ReviewableType: "posts", ReviewCount: 0},

	// 帖子6: 基岩版红石电路和Java版的区别 (我的世界基岩版)
	{ID: 21, UpdatedAt: time.Now().Add(21 * time.Hour), Content: "终于有人总结这个了！基岩版红石确实有些不同", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},
	{ID: 22, UpdatedAt: time.Now().Add(22 * time.Hour), Content: "基岩版的红石更稳定一些", Attitude: nil, UserID: 5, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},
	{ID: 23, UpdatedAt: time.Now().Add(23 * time.Hour), Content: "Java版的特性更多，但基岩版性能更好", Attitude: nil, UserID: 3, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},

	// 帖子7: 手机版联机小技巧 (我的世界基岩版)
	{ID: 24, UpdatedAt: time.Now().Add(24 * time.Hour), Content: "手机联机总是卡，有解决办法吗？", Attitude: nil, UserID: 7, ReviewableID: 7, ReviewableType: "posts", ReviewCount: 0},
	{ID: 25, UpdatedAt: time.Now().Add(25 * time.Hour), Content: "网络优化的建议很实用", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 7, ReviewableType: "posts", ReviewCount: 0},
	{ID: 26, UpdatedAt: time.Now().Add(26 * time.Hour), Content: "iOS和安卓互联还是有问题", Attitude: lo.ToPtr(false), UserID: 8, ReviewableID: 7, ReviewableType: "posts", ReviewCount: 0},

	// 帖子8: 新手必知的生存要点 (饥荒联机版)
	{ID: 27, UpdatedAt: time.Now().Add(27 * time.Hour), Content: "饥荒新手表示很有帮助！", Attitude: lo.ToPtr(true), UserID: 4, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 28, UpdatedAt: time.Now().Add(28 * time.Hour), Content: "冬天准备工作讲得很详细", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 29, UpdatedAt: time.Now().Add(29 * time.Hour), Content: "还是经常被狗咬死...", Attitude: nil, UserID: 1, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 30, UpdatedAt: time.Now().Add(30 * time.Hour), Content: "建议增加一些角色推荐", Attitude: nil, UserID: 3, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},

	// 帖子9: 第一年秋季种植推荐 (星露谷物语)
	{ID: 31, UpdatedAt: time.Now().Add(31 * time.Hour), Content: "南瓜确实是秋季最赚钱的作物", Attitude: lo.ToPtr(true), UserID: 5, ReviewableID: 9, ReviewableType: "posts", ReviewCount: 0},
	{ID: 32, UpdatedAt: time.Now().Add(32 * time.Hour), Content: "甜宝石莓值得种吗？", Attitude: nil, UserID: 7, ReviewableID: 9, ReviewableType: "posts", ReviewCount: 0},
	{ID: 33, UpdatedAt: time.Now().Add(33 * time.Hour), Content: "小麦也不错，可以做面包", Attitude: nil, UserID: 8, ReviewableID: 9, ReviewableType: "posts", ReviewCount: 0},

	// 帖子10: 今日任务快速完成攻略 (光·遇)
	{ID: 34, UpdatedAt: time.Now().Add(34 * time.Hour), Content: "光遇的日常任务确实需要技巧", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 10, ReviewableType: "posts", ReviewCount: 0},
	{ID: 35, UpdatedAt: time.Now().Add(35 * time.Hour), Content: "收集光之翼的路线很棒！", Attitude: lo.ToPtr(true), UserID: 4, ReviewableID: 10, ReviewableType: "posts", ReviewCount: 0},
	{ID: 36, UpdatedAt: time.Now().Add(36 * time.Hour), Content: "季节任务有更新吗？", Attitude: nil, UserID: 6, ReviewableID: 10, ReviewableType: "posts", ReviewCount: 0},

	// 帖子11-19的评论
	{ID: 37, UpdatedAt: time.Now().Add(37 * time.Hour), Content: "我这边加载也很慢，应该是服务器问题", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 11, ReviewableType: "posts", ReviewCount: 0},
	{ID: 38, UpdatedAt: time.Now().Add(38 * time.Hour), Content: "夜间模式确实很需要，眼睛更舒服", Attitude: lo.ToPtr(true), UserID: 3, ReviewableID: 12, ReviewableType: "posts", ReviewCount: 0},
	{ID: 39, UpdatedAt: time.Now().Add(39 * time.Hour), Content: "举报流程说明很清楚", Attitude: lo.ToPtr(true), UserID: 5, ReviewableID: 13, ReviewableType: "posts", ReviewCount: 0},
	{ID: 40, UpdatedAt: time.Now().Add(40 * time.Hour), Content: "确实有用户在恶意刷屏", Attitude: lo.ToPtr(true), UserID: 7, ReviewableID: 14, ReviewableType: "posts", ReviewCount: 0},
	{ID: 41, UpdatedAt: time.Now().Add(41 * time.Hour), Content: "我想加入！什么时间在线？", Attitude: lo.ToPtr(true), UserID: 8, ReviewableID: 15, ReviewableType: "posts", ReviewCount: 0},
	{ID: 42, UpdatedAt: time.Now().Add(42 * time.Hour), Content: "最近在玩艾尔登法环", Attitude: nil, UserID: 2, ReviewableID: 16, ReviewableType: "posts", ReviewCount: 0},
	{ID: 43, UpdatedAt: time.Now().Add(43 * time.Hour), Content: "确实很热，在家吹空调玩游戏最舒服了", Attitude: nil, UserID: 4, ReviewableID: 17, ReviewableType: "posts", ReviewCount: 0},
	{ID: 44, UpdatedAt: time.Now().Add(44 * time.Hour), Content: "1.21的新生物很有趣", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 18, ReviewableType: "posts", ReviewCount: 0},
	{ID: 45, UpdatedAt: time.Now().Add(45 * time.Hour), Content: "现代建筑风格很棒！", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 19, ReviewableType: "posts", ReviewCount: 0},

	// 对评论的评论 (reviews) - 楼中楼

	// 回复评论1
	{ID: 46, UpdatedAt: time.Now().Add(46 * time.Hour), Content: "是的，这个社区的氛围真的很好", Attitude: lo.ToPtr(true), UserID: 3, ReviewableID: 1, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 47, UpdatedAt: time.Now().Add(47 * time.Hour), Content: "欢迎新朋友！", Attitude: nil, UserID: 1, ReviewableID: 1, ReviewableType: "reviews", ReviewCount: 0},

	// 回复评论5
	{ID: 48, UpdatedAt: time.Now().Add(48 * time.Hour), Content: "我可以分享一些建筑技巧", Attitude: nil, UserID: 2, ReviewableID: 5, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 49, UpdatedAt: time.Now().Add(49 * time.Hour), Content: "求私信交流！", Attitude: lo.ToPtr(true), UserID: 7, ReviewableID: 5, ReviewableType: "reviews", ReviewCount: 0},

	// 回复评论10
	{ID: 50, UpdatedAt: time.Now().Add(50 * time.Hour), Content: "我也觉得太复杂了，新手看不懂", Attitude: lo.ToPtr(true), UserID: 8, ReviewableID: 10, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 51, UpdatedAt: time.Now().Add(51 * time.Hour), Content: "可以从简单的红石电路开始学习", Attitude: nil, UserID: 5, ReviewableID: 10, ReviewableType: "reviews", ReviewCount: 0},

	// 回复评论14
	{ID: 52, UpdatedAt: time.Now().Add(52 * time.Hour), Content: "恭喜通关！我还在死循环中", Attitude: nil, UserID: 6, ReviewableID: 14, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 53, UpdatedAt: time.Now().Add(53 * time.Hour), Content: "多练习就好了，我也死了很多次", Attitude: nil, UserID: 4, ReviewableID: 14, ReviewableType: "reviews", ReviewCount: 0},

	// 回复评论21
	{ID: 54, UpdatedAt: time.Now().Add(54 * time.Hour), Content: "基岩版和Java版确实差别挺大", Attitude: lo.ToPtr(true), UserID: 3, ReviewableID: 21, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 55, UpdatedAt: time.Now().Add(55 * time.Hour), Content: "我更喜欢Java版的模组支持", Attitude: nil, UserID: 7, ReviewableID: 21, ReviewableType: "reviews", ReviewCount: 0},

	// 回复评论24
	{ID: 56, UpdatedAt: time.Now().Add(56 * time.Hour), Content: "可能是网络问题，试试换个wifi", Attitude: nil, UserID: 1, ReviewableID: 24, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 57, UpdatedAt: time.Now().Add(57 * time.Hour), Content: "手机性能也很重要", Attitude: nil, UserID: 2, ReviewableID: 24, ReviewableType: "reviews", ReviewCount: 0},

	// 回复评论27
	{ID: 58, UpdatedAt: time.Now().Add(58 * time.Hour), Content: "饥荒确实需要很多技巧", Attitude: lo.ToPtr(true), UserID: 8, ReviewableID: 27, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 59, UpdatedAt: time.Now().Add(59 * time.Hour), Content: "推荐多看攻略视频", Attitude: nil, UserID: 5, ReviewableID: 27, ReviewableType: "reviews", ReviewCount: 0},

	// 更多评论数据以增加测试覆盖率
	{ID: 60, UpdatedAt: time.Now().Add(60 * time.Hour), Content: "游戏社区就是要互相帮助", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},
	{ID: 61, UpdatedAt: time.Now().Add(61 * time.Hour), Content: "这个帖子质量很高", Attitude: lo.ToPtr(true), UserID: 4, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 62, UpdatedAt: time.Now().Add(62 * time.Hour), Content: "有没有其他boss的攻略？", Attitude: nil, UserID: 1, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 63, UpdatedAt: time.Now().Add(63 * time.Hour), Content: "克苏鲁之眼算是入门boss了", Attitude: nil, UserID: 3, ReviewableID: 5, ReviewableType: "posts", ReviewCount: 0},
	{ID: 64, UpdatedAt: time.Now().Add(64 * time.Hour), Content: "基岩版确实有些特殊机制", Attitude: nil, UserID: 7, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},
	{ID: 65, UpdatedAt: time.Now().Add(65 * time.Hour), Content: "手机玩游戏还是有局限性", Attitude: lo.ToPtr(false), UserID: 5, ReviewableID: 7, ReviewableType: "posts", ReviewCount: 0},
	{ID: 66, UpdatedAt: time.Now().Add(66 * time.Hour), Content: "饥荒联机版比单机版有趣多了", Attitude: lo.ToPtr(true), UserID: 8, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 67, UpdatedAt: time.Now().Add(67 * time.Hour), Content: "星露谷物语真的很治愈", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 9, ReviewableType: "posts", ReviewCount: 0},
	{ID: 68, UpdatedAt: time.Now().Add(68 * time.Hour), Content: "光遇的画面很美", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 10, ReviewableType: "posts", ReviewCount: 0},

	// 一些反对态度的评论
	{ID: 69, UpdatedAt: time.Now().Add(69 * time.Hour), Content: "这个建议不太实用", Attitude: lo.ToPtr(false), UserID: 4, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 70, UpdatedAt: time.Now().Add(70 * time.Hour), Content: "攻略有些过时了", Attitude: lo.ToPtr(false), UserID: 1, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 71, UpdatedAt: time.Now().Add(71 * time.Hour), Content: "这个方法效率不高", Attitude: lo.ToPtr(false), UserID: 3, ReviewableID: 3, ReviewableType: "posts", ReviewCount: 0},
	{ID: 72, UpdatedAt: time.Now().Add(72 * time.Hour), Content: "不太同意这个观点", Attitude: lo.ToPtr(false), UserID: 7, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},
	{ID: 73, UpdatedAt: time.Now().Add(73 * time.Hour), Content: "网络问题确实影响体验", Attitude: lo.ToPtr(false), UserID: 5, ReviewableID: 11, ReviewableType: "posts", ReviewCount: 0},

	// 更多楼中楼评论
	{ID: 74, UpdatedAt: time.Now().Add(74 * time.Hour), Content: "同意楼上的看法", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 69, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 75, UpdatedAt: time.Now().Add(75 * time.Hour), Content: "每个人看法不同吧", Attitude: nil, UserID: 8, ReviewableID: 70, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 76, UpdatedAt: time.Now().Add(76 * time.Hour), Content: "有更好的方法吗？", Attitude: nil, UserID: 6, ReviewableID: 71, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 77, UpdatedAt: time.Now().Add(77 * time.Hour), Content: "理性讨论比较好", Attitude: nil, UserID: 4, ReviewableID: 72, ReviewableType: "reviews", ReviewCount: 0},
	{ID: 78, UpdatedAt: time.Now().Add(78 * time.Hour), Content: "希望官方能修复这个问题", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 73, ReviewableType: "reviews", ReviewCount: 0},

	// 补充更多不同类型的评论内容
	{ID: 79, UpdatedAt: time.Now().Add(79 * time.Hour), Content: "大佬能出个详细教程吗？", Attitude: nil, UserID: 3, ReviewableID: 2, ReviewableType: "posts", ReviewCount: 0},
	{ID: 80, UpdatedAt: time.Now().Add(80 * time.Hour), Content: "材料清单能分享一下吗？", Attitude: nil, UserID: 5, ReviewableID: 3, ReviewableType: "posts", ReviewCount: 0},
	{ID: 81, UpdatedAt: time.Now().Add(81 * time.Hour), Content: "这个装备搭配很合理", Attitude: lo.ToPtr(true), UserID: 7, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 82, UpdatedAt: time.Now().Add(82 * time.Hour), Content: "新手第一次打败boss就是这种感觉", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 5, ReviewableType: "posts", ReviewCount: 0},
	{ID: 83, UpdatedAt: time.Now().Add(83 * time.Hour), Content: "两个版本各有优势", Attitude: nil, UserID: 8, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},
	{ID: 84, UpdatedAt: time.Now().Add(84 * time.Hour), Content: "联机确实是手机版的痛点", Attitude: lo.ToPtr(false), UserID: 4, ReviewableID: 7, ReviewableType: "posts", ReviewCount: 0},
	{ID: 85, UpdatedAt: time.Now().Add(85 * time.Hour), Content: "温蒂是新手最友好的角色", Attitude: nil, UserID: 6, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 86, UpdatedAt: time.Now().Add(86 * time.Hour), Content: "蓝莓也是不错的选择", Attitude: nil, UserID: 1, ReviewableID: 9, ReviewableType: "posts", ReviewCount: 0},
	{ID: 87, UpdatedAt: time.Now().Add(87 * time.Hour), Content: "光遇的社交系统很有趣", Attitude: lo.ToPtr(true), UserID: 3, ReviewableID: 10, ReviewableType: "posts", ReviewCount: 0},

	// 最后一批评论，确保数据充分
	{ID: 88, UpdatedAt: time.Now().Add(88 * time.Hour), Content: "支持楼主！", Attitude: lo.ToPtr(true), UserID: 5, ReviewableID: 1, ReviewableType: "posts", ReviewCount: 0},
	{ID: 89, UpdatedAt: time.Now().Add(89 * time.Hour), Content: "收藏了，很有用", Attitude: lo.ToPtr(true), UserID: 7, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 90, UpdatedAt: time.Now().Add(90 * time.Hour), Content: "期待更多这样的内容", Attitude: lo.ToPtr(true), UserID: 2, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 91, UpdatedAt: time.Now().Add(91 * time.Hour), Content: "感谢分享经验", Attitude: lo.ToPtr(true), UserID: 4, ReviewableID: 4, ReviewableType: "posts", ReviewCount: 0},
	{ID: 92, UpdatedAt: time.Now().Add(92 * time.Hour), Content: "简单易懂，赞一个", Attitude: lo.ToPtr(true), UserID: 6, ReviewableID: 5, ReviewableType: "posts", ReviewCount: 0},
	{ID: 93, UpdatedAt: time.Now().Add(93 * time.Hour), Content: "涨知识了", Attitude: lo.ToPtr(true), UserID: 8, ReviewableID: 6, ReviewableType: "posts", ReviewCount: 0},
	{ID: 94, UpdatedAt: time.Now().Add(94 * time.Hour), Content: "实用的小技巧", Attitude: lo.ToPtr(true), UserID: 1, ReviewableID: 7, ReviewableType: "posts", ReviewCount: 0},
	{ID: 95, UpdatedAt: time.Now().Add(95 * time.Hour), Content: "新手必看系列", Attitude: lo.ToPtr(true), UserID: 3, ReviewableID: 8, ReviewableType: "posts", ReviewCount: 0},
	{ID: 96, UpdatedAt: time.Now().Add(96 * time.Hour), Content: "种植推荐很实用", Attitude: lo.ToPtr(true), UserID: 5, ReviewableID: 9, ReviewableType: "posts", ReviewCount: 0},
	{ID: 97, UpdatedAt: time.Now().Add(97 * time.Hour), Content: "任务攻略很详细", Attitude: lo.ToPtr(true), UserID: 7, ReviewableID: 10, ReviewableType: "posts", ReviewCount: 0},
}
