package model

import (
	"gorm.io/gorm"
	"time"
)

/*  mysql  */

type User struct {
	gorm.Model
	Username       string  `gorm:"size:255;unique"` // 用户名
	Nickname       *string `gorm:"size:255"`        // 昵称
	Email          string  `gorm:"size:255;unique"` // 邮箱
	Password       string  `gorm:"size:255"`        // 密码
	Age            *int    // 年龄
	Gender         *string // 性别
	Occupation     *string `gorm:"size:100"`              // 职业
	EducationLevel *string `gorm:"size:100"`              // 教育水平
	AvatarUrl      *string `gorm:"default:null"`          // 头像
	BackgroundUrl  *string `gorm:"default:null"`          // 背景图片地址
	Country        *string `gorm:"default:null"`          // 国家
	City           *string `gorm:"default:null"`          // 城市
	Address        *string `gorm:"default:null"`          // 地址
	AboutMe        *string `gorm:"default:null"`          // 个人介绍
	Roles          []Role  `gorm:"many2many:user_roles;"` // 身份
}

// Role represents a role in the system that groups a set of permissions.
type Role struct {
	gorm.Model
	Name        string       // 身份名称
	Permissions []Permission `gorm:"many2many:role_permissions;"` // 权限
}

// Permission represents a specific action or resource that can be accessed.
type Permission struct {
	gorm.Model
	Name        string // 权限名称
	Description string // 权限描述
}

type Device struct {
	gorm.Model
	UserID          uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"` // 用户ID
	DeviceType      *string `gorm:"size:100"`                                   // 设备类型
	OperatingSystem *string `gorm:"size:100"`                                   // 操作系统
	Browser         *string `gorm:"size:100"`                                   // 浏览器
}

type Detail struct {
	gorm.Model
	Title          string  `gorm:"size:255"`                                        // 标题
	Description    *string `gorm:"type:text"`                                       // 描述
	Categories     uint    `gorm:"index"`                                           // 分类
	CoverImageUrl  *string `gorm:"default:null"`                                    // 封面地址
	Director       *string `gorm:"default:null"`                                    // 导演
	Scriptwriter   *string `gorm:"default:null"`                                    // 编剧
	Actors         *string `gorm:"type:text"`                                       // 演员，e.g. "A,B,C,D"
	CurrentEpisode int     `gorm:"default:1"`                                       // 最新更新的集数
	TotalEpisodes  int     `gorm:"default:1"`                                       // 总集数
	RegionID       uint    `gorm:"index"`                                           // 地区ID
	Year           *int    `gorm:"default:null"`                                    // 年份
	ReleaseTime    *string `gorm:"default:null"`                                    // 上映时间（电影专用）
	Tags           []Tag   `gorm:"many2many:detail_tags;"`                          // 标签
	Videos         []Video `gorm:"foreignKey:DetailID;constraint:OnDelete:CASCADE"` // 关联多个集数，级联删除
	Remark         *string `gorm:"type:text"`                                       // 备注信息
}

type Category struct {
	gorm.Model
	Name  string `gorm:"size:100"` // 分类名字
	Order int    // 排序
}

type Video struct {
	gorm.Model
	DetailID      uint       `gorm:"index;not null;constraint:OnDelete:CASCADE"`     // 所属
	Title         string     `gorm:"size:255"`                                       // 标题
	Description   *string    `gorm:"type:text"`                                      // 描述
	CoverImageUrl *string    `gorm:"default:null"`                                   // 封面地址
	VideoURLs     []VideoURL `gorm:"foreignKey:VideoID;constraint:OnDelete:CASCADE"` // 视频URL（多线路）
	Tags          []Tag      `gorm:"many2many:video_tags;"`                          // 标签
	Remark        *string    `gorm:"type:text"`                                      // 备注信息
}

type VideoURL struct {
	gorm.Model
	VideoID  uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"` // 关联的集数ID
	Order    int     // 排序字段，数字越小排序越靠前
	URL      string  `gorm:"size:255"` // 视频的URL
	Source   *string // 视频的来源
	Quality  *string // 视频的质量（如 720p, 1080p）
	Language *string // 视频的语言（如 "中文", "English")
	Subtitle *string // 字幕信息（如 "内嵌字幕", "无字幕")
}

type Region struct {
	gorm.Model
	Name  string `gorm:"size:100"` // 地区名字
	Order int    // 排序
}

type Carousel struct {
	gorm.Model
	ImageURL string `gorm:"size:1024"` // 轮播图图片URL
	Order    int    // 排序字段，数字越小排序越靠前
	VideoID  uint   `gorm:"index"`     // 关联的视频ID
	Link     string `gorm:"size:1024"` // 链接
	Remark   string `gorm:"type:text"` // 备注信息，使用text类型以存储较长的文本
}

type Tag struct {
	gorm.Model
	TagName string `gorm:"size:100"` // 标签
}

type VideoTag struct {
	VideoID uint `gorm:"primaryKey;autoIncrement:false"` // 视频标签
	TagID   uint `gorm:"primaryKey;autoIncrement:false"` // 标签ID
}

type DetailTag struct {
	DetailID uint `gorm:"primaryKey;autoIncrement:false"` // 系列ID
	TagID    uint `gorm:"primaryKey;autoIncrement:false"` // 标签ID
}

type Advertisement struct {
	gorm.Model
	Advertiser string    `gorm:"size:255"`  // 广告商
	Content    string    `gorm:"type:text"` // 广告内容
	StartDate  time.Time // 广告开始时间
	EndDate    time.Time // 广告结束时间
	IsActive   bool      `gorm:"default:false"` // 广告是否激活
}
