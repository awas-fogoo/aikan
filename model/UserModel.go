package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username      string `gorm:"not null;uniqueIndex"`
	Nickname      string `gorm:"not null"`
	Email         string `gorm:"not null;index"`
	AvatarUrl     string `gorm:"default:'unknown'"` // 头像，默认值为unknown
	BackgroundUrl string `gorm:"default:'unknown'"` // 背景图片，默认值为unknown
	Gender        string `gorm:"default:'unknown'"` // 性别，默认值为unknown
	Age           uint   `gorm:"default:0"`         // 年龄，默认值为0
	Country       string `gorm:"default:'unknown'"` // 国家，默认值为unknown
	City          string `gorm:"default:'unknown'"` // 城市，默认值为unknown
	Address       string `gorm:"default:'unknown'"` // 地址，默认值为unknown
	AboutMe       string `gorm:"default:'unknown'"` // 关于我，默认值为unknown
	Videos        []Video
	Views         []Video
	Likes         []Video
	Collections   []Video
	Following     []*User `gorm:"many2many:user_followings;association_jointable_foreignkey:follower_id"`
	Followers     []*User `gorm:"many2many:user_followers;association_jointable_foreignkey:following_id"`
	Comments      []Comment
	Roles         []Role       `gorm:"many2many:user_roles;"`
	Permissions   []Permission `gorm:"many2many:user_permissions;"`
	Auth          Auth
}

// Auth 定义认证模型
type Auth struct {
	ID       uint `gorm:"primary_key"`
	UserID   uint `gorm:"uniqueIndex"`
	Password string
}

type UserLike struct {
	gorm.Model
	UserID  uint `gorm:"unique_index:user_id_video_like_id"`
	VideoID uint `gorm:"unique_index:user_id_video_like_id"`
}

type UserCollection struct {
	gorm.Model
	UserID  uint `gorm:"unique_index:user_id_video_like_id"`
	VideoID uint `gorm:"unique_index:user_id_video_like_id"`
}

type Video struct {
	gorm.Model
	Title       string   `gorm:"not null"`
	Description string   `gorm:"not null;type:varchar(1000);"`
	Url         string   `gorm:"not null"`
	CoverUrl    string   `gorm:"not null"`
	Views       uint     `gorm:"default:0"`
	Likes       uint     `gorm:"default:0"`
	Collections uint     `gorm:"default:0"`
	Duration    float64  `gorm:"default:0"`
	PartitionID uint     `gorm:"default:0"`
	Review      uint     `gorm:"default:0"`
	Weights     float32  `gorm:"default:0"`
	Quality     string   `gorm:"not null"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	UserID      uint     `gorm:"not null"`
	User        User     `gorm:"foreignKey:UserID"`
	Tag         []Tag
}

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

// VideoTag 中间表
type VideoTag struct {
	gorm.Model
	VideoID uint
	TagID   uint
}

type Danmaku struct {
	gorm.Model
	VideoID uint   `gorm:"not null;index"`
	Content string `gorm:"not null"`
	Color   string `gorm:"not null"`
	Time    uint64 `gorm:"not null"`
	Type    uint8  `gorm:"default:0"` //类型0滚动;1顶部;2底部
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"type:longtext;not null"`
	Description string `gorm:"not null"`
	Videos      []Video
}

type Comment struct {
	gorm.Model
	Content  string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
	VideoID  uint   `gorm:"not null"`
	ParentID *uint
	Children []Comment `gorm:"foreignkey:ParentID"`
}

type CommentRelation struct {
	gorm.Model
	AncestorID   uint //祖先节点ID
	DescendantID uint //后代节点ID
	Level        uint //评论层级
}

type Role struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string
	Permissions []Permission `gorm:"many2many:role_permission;"`
}

type Permission struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string
	Roles       []Role `gorm:"many2many:role_permission;"`
}

type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

type SearchRecord struct {
	gorm.Model
	UserID   uint   // 用户ID
	Keyword  string // 搜索关键词
	Location string // 搜索地点
}
