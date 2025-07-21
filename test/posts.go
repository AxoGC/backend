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
		ReviewCount: 5,
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
		ReviewCount: 7,
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
		ReviewCount: 6,
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
		ReviewCount: 8,
	},

	// 泰拉瑞亚帖子
	{
		ID:          5,
		CreatedAt:   time.Date(2024, 7, 14, 11, 10, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 14, 11, 10, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "克苏鲁之眼简单打法分享",
		Slug:        "eye-of-cthulhu-guide",
		ForumID:     5,
		Content:     "新手向：克苏鲁之眼攻略\n\n准备物品：\n- 银/钨弓 + 燧石箭（至少200发）\n- 轻便靴（提高移动速度）\n- 再生药水、速度药水\n- 血量不少于200\n\n战斗场地：\n在地面建造一个长平台，两端各放一个篝火增加生命恢复。\n\n打法：\n1. 第一阶段：保持距离，左右移动射箭\n2. 第二阶段（血量50%以下）：BOSS会冲刺，注意预判方向\n3. 利用钩爪或平台快速改变方向\n\n掉落物品包括恶魔矿石和腐化种子，是进入中期的重要材料！\n\n祝各位萌新早日击败第一个BOSS！",
		Markdown:    false,
		UserID:      2,
		ReviewCount: 6,
	},

	// 我的世界基岩版帖子
	{
		ID:          6,
		CreatedAt:   time.Date(2024, 7, 13, 13, 25, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 13, 13, 25, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "基岩版红石电路和Java版的区别",
		Slug:        "bedrock-vs-java-redstone",
		ForumID:     2,
		Content:     "经常看到有人问基岩版和Java版红石的区别，这里总结一下：\n\n主要区别：\n1. 准连通性：Java版支持，基岩版不支持\n2. 活塞推拉方块：基岩版更严格，有些Java版能推的方块基岩版推不动\n3. 红石粉更新顺序：两个版本的更新顺序不同，可能导致电路表现差异\n4. 0刻脉冲：基岩版对0刻脉冲的处理更加严格\n\n建议：\n- 如果从Java版转到基岩版，需要重新测试红石电路\n- 基岩版适合做简单实用的红石装置\n- 复杂的红石计算机更适合在Java版制作\n\n有具体问题的朋友可以在评论区讨论！",
		Markdown:    false,
		UserID:      4,
		ReviewCount: 7,
	},
	{
		ID:          7,
		CreatedAt:   time.Date(2024, 7, 16, 19, 40, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 16, 19, 40, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "手机版联机小技巧",
		Slug:        "mobile-multiplayer-tips",
		ForumID:     2,
		Content:     "手机版MC联机经验分享：\n\n1. 网络连接\n- 确保所有玩家连接同一个WiFi网络\n- 或者使用Xbox Live账号进行联机\n\n2. 世界设置\n- 创建世界时记得开启\"对朋友可\"\n- 设置合适的游戏模式和难度\n\n3. 常见问题解决\n- 看不到朋友的世界：检查是否添加为Xbox好友\n- 连接超时：尝试重启游戏或切换网络\n- 存档同步：定期备份世界存档\n\n4. 联机礼仪\n- 不要随意破坏别人的建筑\n- 共同决定重要的建设项目\n- 合理分配资源\n\n希望大家都能愉快联机！",
		Markdown:    false,
		UserID:      6,
		ReviewCount: 6,
	},

	// 饥荒联机版帖子
	{
		ID:          8,
		CreatedAt:   time.Date(2024, 7, 11, 16, 55, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 11, 16, 55, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "新手必知的生存要点",
		Slug:        "dst-survival-guide",
		ForumID:     4,
		Content:     "饥荒新手生存指南：\n\n第一天：\n- 捡树枝、草、燧石\n- 做斧头砍树，做镐挖石头\n- 找到浆果丛和胡萝卜\n- 在天黑前做篝火\n\n前三天重点：\n- 探索地图，寻找重要资源点\n- 收集足够食物（浆果、胡萝卜、种子）\n- 做背包增加库存\n- 找到适合建家的地方\n\n建家选址：\n- 靠近浆果丛和草原\n- 远离猪人村和蜘蛛巢\n- 有石头矿和树木\n\n重要提醒：\n- 永远不要让篝火熄灭！\n- 理智值很重要，多做花环\n- 不要贪心，活着最重要\n\n祝新手们都能度过第一个冬天！",
		Markdown:    false,
		UserID:      8,
		ReviewCount: 11,
	},

	// 星露谷物语帖子
	{
		ID:          9,
		CreatedAt:   time.Date(2024, 7, 17, 10, 20, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 17, 10, 20, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "第一年秋季种植推荐",
		Slug:        "fall-year1-crops",
		ForumID:     6,
		Content:     "第一年秋季作物种植攻略：\n\n最佳选择：\n1. 蔓越莓（种子240G）\n   - 收获周期：5天，之后每2天收获一次\n   - 利润：整个秋季约1920G/格\n\n2. 南瓜（种子100G）\n   - 收获周期：13天\n   - 制作南瓜汁利润更高\n\n3. 甜玉米（种子150G）\n   - 夏秋两季作物，如果夏天种了秋天继续收获\n\n种植建议：\n- 重点种植蔓越莓，性价比最高\n- 留一些地方种南瓜做Bundle\n- 记得做喷水器提高效率\n- 别忘了升级洒水壶\n\n秋季也是采集的好时机，松露、黑莓都很值钱！\n\n大家第一年秋季都种什么呢？",
		Markdown:    false,
		UserID:      1,
		ReviewCount: 6,
	},

	// 光·遇帖子
	{
		ID:          10,
		CreatedAt:   time.Date(2024, 7, 19, 8, 30, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 19, 8, 30, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "今日任务快速完成攻略",
		Slug:        "daily-quests-guide",
		ForumID:     7,
		Content:     "每日任务效率完成法：\n\n常见任务类型：\n1. 收集蜡烛/光芒\n   - 优先去人多的地图\n   - 云野、雨林效率最高\n\n2. 拯救精灵\n   - 记住固定刷新点\n   - 可以和朋友轮流带路\n\n3. 重温先祖记忆\n   - 提前记好先祖位置\n   - 某些先祖需要多人配合\n\n4. 社交任务\n   - 加入活跃的群聊\n   - 主动向其他玩家表达动作\n\n小贴士：\n- 任务可以跨地图完成\n- 和朋友组队效率更高\n- 每周日任务会重置\n- 完成任务获得的季节蜡烛很珍贵\n\n希望小光们都能轻松完成每日任务！✨",
		Markdown:    false,
		UserID:      3,
		ReviewCount: 6,
	},

	// BUG反馈帖子
	{
		ID:          11,
		CreatedAt:   time.Date(2024, 7, 20, 12, 15, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 20, 12, 15, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "手机端页面加载速度慢",
		Slug:        "mobile-loading-slow",
		ForumID:     8,
		Content:     "反馈一个问题：\n\n设备：小米12 Pro\n浏览器：Chrome\n网络：4G/WiFi都试过了\n\n问题描述：\n手机端打开论坛首页需要等待10-15秒才能完全加载完成，点击帖子也比较慢。电脑端正常。\n\n其他用户有遇到类似问题吗？\n\n建议：\n- 考虑对手机端做一些优化\n- 图片可以做懒加载\n- 减少首页加载的内容\n\n整体使用体验还是很不错的，就是速度有点慢 :)",
		Markdown:    false,
		UserID:      5,
		ReviewCount: 2,
	},
	{
		ID:          12,
		CreatedAt:   time.Date(2024, 7, 18, 22, 45, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 18, 22, 45, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "建议增加夜间模式",
		Slug:        "add-dark-mode",
		ForumID:     8,
		Content:     "希望能加个夜间模式功能！\n\n原因：\n- 晚上使用白色背景比较刺眼\n- 夜间模式更省电\n- 现在很多网站都有这个功能了\n\n建议的实现方式：\n- 右上角加个切换按钮\n- 可以跟随系统设置自动切换\n- 记住用户的选择偏好\n\n配色建议：\n- 深色背景 + 浅色文字\n- 保持原有的强调色不变\n- 适当降低对比度保护眼睛\n\n期待这个功能的实现！🌙",
		Markdown:    false,
		UserID:      7,
		ReviewCount: 1,
	},

	// 举报帖子
	{
		ID:          13,
		CreatedAt:   time.Date(2024, 7, 16, 14, 20, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 16, 14, 20, 0, 0, time.UTC),
		Pinned:      true,
		Title:       "举报规范和处理流程说明",
		Slug:        "report-guidelines",
		ForumID:     9,
		Content:     "# 举报规范说明\n\n## 可举报的行为类型\n\n1. **恶意灌水**：无意义的重复发帖或回复\n2. **人身攻击**：辱骂、诽谤其他用户\n3. **广告推广**：发布与论坛主题无关的广告内容\n4. **违规内容**：色情、暴力、政治敏感内容\n5. **恶意刷屏**：短时间内大量发帖影响正常讨论\n\n## 举报格式\n\n请按照以下格式提供举报信息：\n- **举报对象**：用户ID或帖子链接\n- **违规类型**：选择上述类型之一\n- **违规内容**：简要描述具体行为\n- **相关证据**：截图或链接（如有）\n\n## 处理流程\n\n1. 管理员会在24小时内查看举报\n2. 核实违规行为后会采取相应措施\n3. 处理结果会通过站内信告知举报者\n4. 恶意举报也会受到相应处罚\n\n感谢大家共同维护社区环境！",
		Markdown:    true,
		UserID:      1,
		ReviewCount: 1,
	},
	{
		ID:          14,
		CreatedAt:   time.Date(2024, 7, 19, 16, 30, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 19, 16, 30, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "举报用户恶意刷屏",
		Slug:        "report-spam-user",
		ForumID:     9,
		Content:     "举报对象：用户ID 999（示例）\n违规类型：恶意刷屏\n\n违规内容：\n该用户在\"聊天室\"版块短时间内发布了20多个无意义的帖子，内容都是\"顶\"、\"沙发\"、\"路过\"等灌水内容，严重影响了正常的讨论秩序。\n\n时间：2024年7月19日 15:30-16:00\n\n希望管理员能够及时处理，谢谢！\n\n（注：这是示例举报，实际使用时请提供真实的用户信息）",
		Markdown:    false,
		UserID:      4,
		ReviewCount: 1,
	},

	// 更多聊天室帖子
	{
		ID:          15,
		CreatedAt:   time.Date(2024, 7, 19, 21, 10, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 19, 21, 10, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "有没有人一起组队玩游戏？",
		Slug:        "looking-for-gaming-partners",
		ForumID:     1,
		Content:     "最近一个人玩游戏有点无聊，想找几个朋友一起组队！\n\n我常玩的游戏：\n- Minecraft（Java和基岩版都可以）\n- 饥荒联机版\n- 星露谷物语联机\n- 泰拉瑞亚\n\n游戏时间：一般晚上8点后有空\nQQ群：123456789（示例）\n\n希望找到：\n- 有基本游戏素养的朋友\n- 能够友好交流，不会随意破坏\n- 最好有语音，游戏体验更好\n\n欢迎萌新和大佬都来，一起愉快游戏！🎮",
		Markdown:    false,
		UserID:      2,
		ReviewCount: 1,
	},
	{
		ID:          16,
		CreatedAt:   time.Date(2024, 7, 15, 14, 30, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 15, 14, 30, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "大家都在玩什么游戏？",
		Slug:        "what-games-are-you-playing",
		ForumID:     1,
		Content:     "最近游戏荒了，想找点新游戏玩玩。大家最近都在玩什么有意思的游戏吗？求推荐！不限平台和类型～",
		Markdown:    false,
		UserID:      3,
		ReviewCount: 1,
	},
	{
		ID:          17,
		CreatedAt:   time.Date(2024, 7, 18, 20, 15, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 18, 20, 15, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "今天天气好热啊",
		Slug:        "so-hot-today",
		ForumID:     1,
		Content:     "外面40度+的天气，只想在家吹空调玩游戏。大家都是怎么度过这个夏天的呢？有什么消暑的好方法吗？",
		Markdown:    false,
		UserID:      7,
		ReviewCount: 1,
	},

	// 我的世界Java版帖子
	{
		ID:          18,
		CreatedAt:   time.Date(2024, 7, 10, 9, 45, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 16, 16, 20, 0, 0, time.UTC),
		Pinned:      true,
		Title:       "1.21版本新特性总结",
		Slug:        "minecraft-1-21-features",
		ForumID:     3,
		Content:     "# Minecraft 1.21 版本更新内容\n\n## 新增内容\n\n### 新方块\n- **铜灯泡**：可以被红石信号控制的装饰性光源\n- **凝灰岩系列**：包括凝灰岩、磨制凝灰岩、凝灰岩砖等\n- **铜门和铜活板门**：会随时间氧化变色的门类方块\n\n### 新生物群系\n- **试炼密室**：包含试炼刷怪笼的全新地下结构\n- **苍白花园**：神秘的新生物群系\n\n### 游戏机制优化\n- 改进了红石电路的性能\n- 优化了区块加载速度\n- 修复了多项已知bug\n\n欢迎大家分享自己的游戏体验！",
		Markdown:    true,
		UserID:      1,
		ReviewCount: 1,
	},
	{
		ID:          19,
		CreatedAt:   time.Date(2024, 7, 12, 15, 20, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 7, 12, 15, 20, 0, 0, time.UTC),
		Pinned:      false,
		Title:       "分享一下我的现代别墅建筑",
		Slug:        "my-modern-villa-build",
		ForumID:     3,
		Content:     "花了一个月时间建造的现代别墅终于完工了！\n\n特色：\n- 大面积玻璃幕墙设计\n- 无边际泳池\n- 地下车库\n- 屋顶花园\n- 全自动照明系统\n\n主要材料：石英块、玻璃、混凝土、橡木等。整体采用白色调配黑色点缀，简约现代风格。\n\n建造过程中遇到最大的难题是如何让建筑看起来不那么方正，最后通过增加斜面屋顶和不规则阳台解决了这个问题。\n\n有兴趣的朋友可以加服务器来参观，坐标我私聊发给大家！",
		Markdown:    false,
		UserID:      5,
		ReviewCount: 1,
	},
}
