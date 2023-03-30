package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username    string       `gorm:"not null;uniqueIndex"`
	Password    string       `gorm:"not null"`
	Nickname    string       `gorm:"not null"`
	Email       string       `gorm:"not null;index"`
	AvatarUrl   string       `gorm:"not null"`
	Roles       []Role       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
	Videos      []Video
	Likes       []Video `gorm:"many2many:user_likes;association_foreignkey:id;foreignkey:id"`
	Followees   []*User `gorm:"many2many:user_followees;ForeignKey:FollowerID"`
	Followers   []*User `gorm:"many2many:user_followers;ForeignKey:FolloweeID"`
	Comments    []Comment
}

type Video struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Url         string `gorm:"not null"`
	CoverUrl    string `gorm:"not null"`
	Views       int
	Likes       int
	Collections int
	Quality     string `gorm:"not null"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	Comments    []Comment
	Tags        string `gorm:"not null"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Videos      []Video
}

type Comment struct {
	gorm.Model
	Content  string `gorm:"not null"`
	UserID   uint
	VideoID  uint
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
