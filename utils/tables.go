package utils

import (
	"time"
)

// 从上到下执行自动迁移，下者依赖上者
var Tables = []any{

	// A组：不依赖其他组的模型
	// A1组：枚举模型
	new(Role), new(Prop), new(Game),
	// A2组：不依赖的模型
	new(User), new(Guild), new(DocGroup), new(ForumGroup),
	// A3组：只依赖枚举的模型
	new(Good), new(Server),

	// B组：只依赖A组的模型
	new(UserGuild), new(UserFollow), new(Donation), new(UserRole),
	new(Doc), new(Album), new(File), new(Review), new(UserProp), new(Log),
	new(DeepSeekMessage),

	// C组：依赖A组或B组的模型
	new(Online), new(Image), new(Forum),

	// D组：依赖A组或B组或C组的模型
	new(Post),
}

type RoleID string

type Role struct {
	ID          RoleID `gorm:"type:VARCHAR(32);primarykey;comment:ID"`
	Label       string `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Description string `gorm:"type:VARCHAR(255);not null;comment:描述"`
}

const (
	Admin        RoleID = "admin"
	WikiAdmin    RoleID = "wiki_admin"
	BBSAdmin     RoleID = "bbs_admin"
	GalleryAdmin RoleID = "gallery_admin"
	ReviewAdmin  RoleID = "review_admin"
)

type PropID string

type Prop struct {
	ID          PropID `gorm:"primarykey;type:VARCHAR(32);comment:ID"`
	Label       string `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Description string `gorm:"type:VARCHAR(255);not null;comment:描述"`
}

const (
	BlindBox PropID = "blind_box"
)

type GameID string

type Game struct {
	ID          GameID `gorm:"type:VARCHAR(32);primarykey;comment:ID"`
	Label       string `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Description string `gorm:"type:VARCHAR(255);not null;comment:描述"`
}

const (
	MinecraftBedrock   GameID = "minecraft_bedrock"
	MinecraftJava      GameID = "minecraft_java"
	DontStarve         GameID = "dont_starve"
	Terraria           GameID = "terraria"
	StardewValley      GameID = "stardew_valley"
	Palworld           GameID = "palworld"
	ARKSurvivalEvolved GameID = "ark_survival_evolved"
)

type User struct {
	ID             uint           `gorm:"comment:ID"`
	CreatedAt      time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt      time.Time      `gorm:"not null;comment:更新时间"`
	Name           string         `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Exp            uint           `gorm:"not null;comment:经验值"`
	Password       string         `gorm:"type:CHAR(60);not null;comment:密码"`
	Gender         *bool          `gorm:"comment:性别"`
	Profile        string         `gorm:"type:VARCHAR(255);not null;comment:个人介绍"`
	Birthday       *time.Time     `gorm:"comment:生日"`
	Location       string         `gorm:"type:VARCHAR(128);not null;comment:地址"`
	DailyCoin      uint           `gorm:"not null;comment:签到币"`
	HonorCoin      uint           `gorm:"not null;comment:贡献币"`
	Checkin        int64          `gorm:"not null;comment:签到记录"`
	Email          string         `gorm:"type:VARCHAR(128);not null;unique;comment:邮箱"`
	QQ             *string        `gorm:"type:VARCHAR(32);unique;comment:QQ号"`
	MCBEName       *string        `gorm:"type:VARCHAR(32);unique;comment:MCBE用户名"`
	MCJEName       *string        `gorm:"type:VARCHAR(32);unique;comment:MCJE用户名"`
	Setting        map[string]any `gorm:"type:JSON;serializer:json;comment:用户设置"`
	FollowingCount uint           `gorm:"not null;comment:关注数量"`
	FollowerCount  uint           `gorm:"not null;comment:粉丝数量"`
	Following      []User         `gorm:"many2many:user_follows;joinForeignKey:follower_id;joinReferences:following_id"`
	Followers      []User         `gorm:"many2many:user_follows;joinForeignKey:following_id;joinReferences:follower_id"`
	UserRoles      []UserRole
	UserGuild      *UserGuild
}

type Guild struct {
	ID         uint      `gorm:"comment:ID"`
	CreatedAt  time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"not null;comment:更新时间"`
	Name       string    `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Slug       string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	UserCount  uint      `gorm:"not null;comment:公会人数"`
	Profile    string    `gorm:"type:VARCHAR(255);not null;comment:公会介绍"`
	Notice     string    `gorm:"type:TEXT;not null;comment:公会公告"`
	Money      uint      `gorm:"not null;comment:公会资金"`
	UserGuilds []UserGuild
}

type DocGroup struct {
	ID    uint   `gorm:"comment:ID"`
	Label string `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Sort  int    `gorm:"not null;comment:排序"`
	Docs  []Doc
}

type ForumGroup struct {
	ID     uint   `gorm:"comment:ID"`
	Label  string `gorm:"type:VARCHAR(32);not null;unique;comment:标题"`
	Sort   int    `gorm:"not null;comment:排序"`
	Forums []Forum
}

type Good struct {
	ID     uint   `gorm:"comment:ID"`
	PropID PropID `gorm:"type:VARCHAR(32);not null;index;comment:物品"`
	Prop   Prop   `gorm:"constraint:OnDelete:RESTRICT"`
	Label  string `gorm:"type:VARCHAR(32);not null;unique;comment:描述"`
	Count  uint   `gorm:"not null;comment:数量"`
	Price  uint   `gorm:"not null;comment:价格"`
}

type Server struct {
	ID           uint           `gorm:"comment:ID"`
	Slug         string         `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Label        string         `gorm:"type:VARCHAR(32);not null;unique;comment:名称"`
	Path         string         `gorm:"type:VARCHAR(128);not null;unique;comment:路径"`
	Port         uint16         `gorm:"not null;unique;comment:端口"`
	Description  string         `gorm:"type:VARCHAR(255);not null;comment:简介"`
	GameID       GameID         `gorm:"type:VARCHAR(32);index;not null;comment:游戏类型ID"`
	Game         Game           `gorm:"constraint:OnDelete:RESTRICT"`
	BackupEnable bool           `gorm:"not null;comment:启用备份"`
	BackupPath   string         `gorm:"type:VARCHAR(128);not null;unique;comment:备份路径"`
	BackupCron   string         `gorm:"type:VARCHAR(32);not null;comment:备份频率"`
	BackupLimit  uint           `gorm:"not null;comment:备份数量"`
	Meta         map[string]any `gorm:"type:JSON;serializer:json;comment:元信息"`
}

type UserGuild struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UserID    uint      `gorm:"not null;unique;comment:用户"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	Status    *bool     `gorm:"not null;comment:状态"`
	Admin     bool      `gorm:"not null;comment:管理员"`
	GuildID   uint      `gorm:"not null;index;comment:公会ID"`
	Guild     Guild     `gorm:"constraint:OnDelete:CASCADE"`
}

type UserFollow struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	FollowerID  uint      `gorm:"not null;uniqueIndex:idx_follower_following;comment:关注者"`
	FollowingID uint      `gorm:"not null;uniqueIndex:idx_follower_following;comment:被关注者"`
	Follower    User      `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Following   User      `gorm:"foreignKey:FollowingID;constraint:OnDelete:CASCADE"`
}

type Donation struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Amount    float64   `gorm:"not null;comment:金额"`
	UserID    *uint     `gorm:"index;comment:用户"`
	User      *User     `gorm:"constraint:OnDelete:SET NULL"`
	Message   string    `gorm:"type:VARCHAR(128);not null;comment:留言"`
}

type UserRole struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UserID    uint      `gorm:"index;not null;comment:用户"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	RoleID    RoleID    `gorm:"type:VARCHAR(32);index;not null;comment:权限"`
	Role      Role      `gorm:"constraint:OnDelete:RESTRICT"`
}

type Doc struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	Slug        string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Title       string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	DocGroupID  uint      `gorm:"not null;index;comment:知识库ID"`
	DocGroup    DocGroup  `gorm:"constraint:OnDelete:CASCADE"`
	UserID      uint      `gorm:"not null;index;comment:最后编辑者ID"`
	User        User      `gorm:"constraint:OnDelete:CASCADE"`
	Content     string    `gorm:"type:TEXT;not null;comment:内容"`
	Sort        int       `gorm:"not null;comment:排序"`
	ReviewCount uint      `gorm:"not null;comment:评论数量"`
}

type Album struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	UserID      uint      `gorm:"index;not null;comment:创建者ID"`
	User        User      `gorm:"constraint:OnDelete:CASCADE"`
	GuildID     *uint     `gorm:"index;comment:公会ID"`
	Guild       *Guild    `gorm:"constraint:OnDelete:SET NULL"`
	Slug        string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Label       string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	Profile     string    `gorm:"type:VARCHAR(255);not null;comment:简介"`
	Pinned      bool      `gorm:"not null;comment:是否置顶"`
	Hide        bool      `gorm:"not null;comment:隐藏"`
	Protected   bool      `gorm:"not null;comment:上传保护"`
	ImageCount  uint      `gorm:"not null;comment:图片数量"`
	ReviewCount uint      `gorm:"not null;comment:评论数量"`
}

type File struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	Name        string    `gorm:"type:VARCHAR(128);index;not null;comment:文件名"`
	UserID      uint      `gorm:"index;not null;comment:上传者"`
	User        User      `gorm:"constraint:OnDelete:CASCADE"`
	Description string    `gorm:"type:VARCHAR(255);not null;comment:描述"`
}

type Review struct {
	ID             uint      `gorm:"comment:ID"`
	UpdatedAt      time.Time `gorm:"not null;comment:更新时间"`
	Content        string    `gorm:"type:VARCHAR(255);not null;comment:内容"`
	Attitude       *bool     `gorm:"comment:态度"`
	UserID         uint      `gorm:"index;not null;comment:作者ID"`
	User           User      `gorm:"constraint:OnDelete:CASCADE"`
	ReviewableID   uint      `gorm:"not null;index:idx_reviewable;comment:对象ID"`
	ReviewableType string    `gorm:"type:VARCHAR(32);not null;index:idx_reviewable;comment:对象类型"`
	ReviewCount    uint      `gorm:"not null;comment:评论数量"`
}

type UserProp struct {
	ID        uint      `gorm:"comment:ID"`
	UpdatedAt time.Time `gorm:"not null;comment:更新时间"`
	PropID    PropID    `gorm:"type:VARCHAR(32);not null;index;comment:道具类型ID"`
	Prop      Prop      `gorm:"constraint:OnDelete:RESTRICT"`
	UserID    uint      `gorm:"not null;index;comment:拥有者ID"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	Count     uint      `gorm:"not null;comment:数量"`
}

type Log struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Path      string    `gorm:"type:VARCHAR(32);not null;comment:路径"`
	Method    string    `gorm:"type:VARCHAR(32);not null;comment:请求方法"`
	Status    int       `gorm:"not null;comment:状态码"`
	UserID    *uint     `gorm:"index;comment:用户ID"`
	User      *User     `gorm:"constraint:OnDelete:SET NULL"`
	Message   string    `gorm:"type:VARCHAR(255);not null;comment:错误信息"`
}

type DeepSeekMessage struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	UserID    uint      `gorm:"not null;index;comment:用户"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	Role      string    `gorm:"type:VARCHAR(32);not null;comment:角色"`
	Content   string    `gorm:"type:TEXT;not null;comment:消息"`
}

type Online struct {
	ID       uint      `gorm:"comment:ID"`
	Time     time.Time `gorm:"not null;comment:创建时间"`
	ServerID uint      `gorm:"index;not null;comment:服务器"`
	Server   Server    `gorm:"constraint:OnDelete:CASCADE"`
	Count    uint      `gorm:"not null;comment:在线人数"`
}

type Image struct {
	ID        uint      `gorm:"comment:ID"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Filename  string    `gorm:"type:VARCHAR(64);not null;unique;comment:文件名"`
	Label     string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	Likes     uint      `gorm:"not null;comment:点赞"`
	UserID    uint      `gorm:"not null;index;comment:上传者用户ID"`
	User      User      `gorm:"constraint:OnDelete:CASCADE"`
	AlbumID   uint      `gorm:"not null;index;comment:相册ID"`
	Album     Album     `gorm:"constraint:OnDelete:CASCADE"`
}

type Forum struct {
	ID           uint       `gorm:"comment:ID"`
	ForumGroupID uint       `gorm:"not null;index;comment:论坛组ID"`
	ForumGroup   ForumGroup `gorm:"constraint:OnDelete:CASCADE"`
	Slug         string     `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	Title        string     `gorm:"type:VARCHAR(32);not null;unique;comment:标题"`
	SubTitle     string     `gorm:"type:VARCHAR(32);not null;comment:副标题"`
	Profile      string     `gorm:"type:VARCHAR(255);not null;comment:简介"`
	PostCount    uint       `gorm:"not null;comment:帖子数量"`
	Sort         int        `gorm:"not null;comment:排序"`
	ServerID     *uint      `gorm:"index;comment:服务器"`
	Server       *Server    `gorm:"constraint:OnDelete:SET NULL"`
	Posts        []Post
}

type Post struct {
	ID          uint      `gorm:"comment:ID"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	Pinned      bool      `gorm:"not null;comment:是否置顶"`
	Title       string    `gorm:"type:VARCHAR(32);not null;comment:标题"`
	Slug        string    `gorm:"type:VARCHAR(32);not null;unique;comment:标识"`
	ForumID     uint      `gorm:"not null;index;comment:论坛ID"`
	Forum       Forum     `gorm:"constraint:OnDelete:CASCADE"`
	Content     string    `gorm:"type:TEXT;not null;comment:原内容"`
	Markdown    bool      `gorm:"not null;comment:启用markdown"`
	UserID      uint      `gorm:"not null;index;comment:作者ID"`
	User        User      `gorm:"constraint:OnDelete:CASCADE"`
	ReviewCount uint      `gorm:"not null;comment:评论数量"`
}
