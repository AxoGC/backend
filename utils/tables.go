package utils

import (
	"time"
)

var Tables = []any{
	new(ForumGroup), // 论坛组表
	new(DocGroup),   // 知识库表
	new(Guild),      // 公会表

	new(Forum), // 论坛表
	new(User),  // 用户表
	new(Doc),   // 文档表

	new(Post),     // 帖子表
	new(Album),    // 文件夹表
	new(Image),    // 文件表
	new(Review),   // 评论表
	new(Prop),     // 道具类型表
	new(UserProp), // 道具表
	new(Online),   // 在线记录表
	new(Log),      // 日志表

	new(Tag),
}

type GuildRole string

type User struct {
	ID        uint           `gorm:"comment:ID"`
	CreatedAt time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"not null;comment:更新时间"`
	Name      string         `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Exp       uint           `gorm:"not null;comment:经验值"`
	Level     uint           `gorm:"-"`
	Password  string         `gorm:"type:CHAR(64);not null;comment:密码"`
	Admin     bool           `gorm:"not null;default:false;comment:是否管理员"`
	Gender    *bool          `gorm:"comment:性别"`
	Profile   string         `gorm:"type:VARCHAR(255);not null;comment:个人介绍"`
	Birthday  *time.Time     `gorm:"comment:生日"`
	Location  string         `gorm:"type:VARCHAR(128);not null;comment:地址"`
	DailyCoin uint           `gorm:"not null;comment:签到币"`
	HonorCoin uint           `gorm:"not null;comment:贡献币"`
	Checkin   int64          `gorm:"not null;comment:签到记录"`
	Email     string         `gorm:"type:VARCHAR(128);not null;unique;comment:邮箱"`
	QQ        *string        `gorm:"type:VARCHAR(32);unique;comment:QQ号"`
	MCBEName  *string        `gorm:"type:VARCHAR(32);unique;comment:MCBE用户名"`
	MCJEName  *string        `gorm:"type:VARCHAR(32);unique;comment:MCJE用户名"`
	GuildID   *uint          `gorm:"index;comment:所属公会"`
	Guild     *Guild         `gorm:"constraint:OnDelete:SET NULL"`
	GuildRole GuildRole      `gorm:"type:VARCHAR(32);not null;comment:公会身份"`
	Donation  uint           `gorm:"not null;comment:捐赠数额"`
	Setting   map[string]any `gorm:"type:JSON;serializer:json;not null;comment:用户设置"`
}

type Role struct {
	ID    uint   `gorm:"comment:ID"`
	Label string `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Key   string `gorm:"type:VARCHAR(32);not null;unique;comment:键"`
}

type UserRole struct {
	ID     uint `gorm:"comment:ID"`
	UserID uint `gorm:"index;not null;comment:用户"`
	RoleID uint `gorm:"index;not null;comment:权限"`
}

type DocGroup struct {
	ID    uint   `gorm:"comment:ID"`
	Label string `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Sort  int    `gorm:"not null;comment:排序"`
}

type Doc struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	Slug        string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Title       string    `gorm:"type:VARCHAR(32);uniqueIndex:idx_title_group;not null;comment:标题"`
	DocGroupID  uint      `gorm:"uniqueIndex:idx_title_group;not null;comment:知识库ID"`
	UserID      uint      `gorm:"index;not null;comment:最后编辑者ID"`
	Content     string    `gorm:"type:TEXT;not null;comment:内容"`
	Sort        int       `gorm:"not null;comment:排序"`
	ReviewCount uint      `gorm:"not null;comment:评论数量"`
	Tags        []Tag     `gorm:"many2many:doc_tags;"`
}

type Tag struct {
	ID    uint   `gorm:"id"`
	Label string `gorm:"type:VARCHAR(32);not null;unique;comment:标签"`
	Docs  []Doc  `gorm:"many2many:doc_tags;"`
}

type Online struct {
	ID     uint      `gorm:"comment:ID"`
	Time   time.Time `gorm:"not null;comment:创建时间"`
	Server string    `gorm:"type:VARCHAR(32);not null;index;comment:服务器"`
	Count  int       `gorm:"not null;comment:在线人数"`
}

type Guild struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	Name      string    `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Slug      string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	UserCount uint      `gorm:"not null;comment:公会人数"`
	Profile   string    `gorm:"type:VARCHAR(255);not null;comment:公会介绍"`
	Notice    string    `gorm:"type:TEXT;not null;comment:公会公告"`
	Money     uint      `gorm:"not null;comment:公会资金"`
}

type Album struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	UserID      uint      `gorm:"index;not null;comment:创建者ID"`
	GuildID     *uint     `gorm:"index;comment:公会ID"`
	Slug        string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Label       string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	Profile     string    `gorm:"type:VARCHAR(255);not null;comment:简介"`
	Pinned      bool      `gorm:"not null;comment:是否置顶"`
	Private     bool      `gorm:"not null;comment:保密"`
	Protected   bool      `gorm:"not null;comment:受保护"`
	ImageCount  uint      `gorm:"not null;comment:图片数量"`
	ReviewCount uint      `gorm:"not null;comment:评论数量"`
}

type Image struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Filename  string    `gorm:"type:VARCHAR(64);not null;unique;comment:文件名"`
	Label     string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	Likes     uint      `gorm:"not null;comment:点赞"`
	UserID    uint      `gorm:"index;not null;comment:上传者用户ID"`
	AlbumID   uint      `gorm:"index;not null;comment:相册ID"`
}

type ForumGroup struct {
	ID    uint   `gorm:"comment:ID"`
	Label string `gorm:"type:VARCHAR(32);not null;unique;comment:标题"`
	Sort  int    `gorm:"not null;comment:排序"`
}

type Forum struct {
	ID           uint   `gorm:"comment:ID"`
	ForumGroupID uint   `gorm:"index;not null;comment:论坛组ID"`
	Slug         string `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Title        string `gorm:"type:VARCHAR(32);not null;unique;comment:标题"`
	SubTitle     string `gorm:"type:VARCHAR(32);not null;comment:副标题"`
	Profile      string `gorm:"type:VARCHAR(255);not null;comment:简介"`
	PostCount    uint   `gorm:"not null;comment:帖子数量"`
	Sort         int    `gorm:"not null;comment:排序"`
}

type Post struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	Pinned      bool      `gorm:"not null;comment:是否置顶"`
	Title       string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	Slug        string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	ForumID     uint      `gorm:"index;comment:论坛ID"`
	Content     string    `gorm:"type:TEXT;not null;comment:原内容"`
	Markdown    bool      `gorm:"not null;comment:启用markdown"`
	UserID      uint      `gorm:"index;not null;comment:作者ID"`
	ReviewCount uint      `gorm:"not null;comment:评论数量"`
}

type Review struct {
	ID             uint      `gorm:"comment:ID"`
	UpdatedAt      time.Time `gorm:"not null;comment:更新时间"`
	Content        string    `gorm:"type:VARCHAR(255);not null;comment:内容"`
	Attitude       *bool     `gorm:"comment:态度"`
	UserID         uint      `gorm:"index;not null;comment:作者ID"`
	ReviewableID   uint      `gorm:"index:idx_reviewable;not null;comment:对象ID"`
	ReviewableType string    `gorm:"index:idx_reviewable;not null;type:VARCHAR(32);comment:对象类型"`
	ReviewCount    uint      `gorm:"not null;comment:评论数量"`
}

type UserProp struct {
	ID        uint      `gorm:"comment:ID"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	PropID    uint      `gorm:"index;not null;comment:道具类型ID"`
	UserID    uint      `gorm:"index;not null;comment:拥有者ID"`
	Count     uint      `gorm:"not null;comment:数量"`
}

type Prop struct {
	ID       uint   `gorm:"comment:ID"`
	Label    string `gorm:"type:VARCHAR(32);not null;unique;comment:道具名称"`
	Price    *uint  `gorm:"comment:价格"`
	Profile  string `gorm:"type:VARCHAR(255);not null;comment:道具简介"`
	Function string `gorm:"type:VARCHAR(32);not null;comment:功能"`
	Params   []byte `gorm:"type:JSON;serializer:json;not null;comment:功能参数"`
}

type Log struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Path      string    `gorm:"type:VARCHAR(32);not null;comment:路径"`
	Method    string    `gorm:"type:VARCHAR(32);not null;comment:请求方法"`
	Status    int       `gorm:"not null;comment:状态码"`
	UserID    *uint     `gorm:"index;comment:用户ID"`
	Error     string    `gorm:"type:VARCHAR(255);not null;comment:错误信息"`
}
