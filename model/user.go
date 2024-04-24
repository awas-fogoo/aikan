package model

import (
	"gorm.io/gorm"
	"time"
)

/*  mysql  */

type User struct {
	gorm.Model
	Username       string  `gorm:"size:255;unique"`
	Nickname       *string `gorm:"size:255"`
	Email          string  `gorm:"size:255;unique"`
	Password       string  `gorm:"size:255"`
	Age            *int
	Gender         *string
	Occupation     *string `gorm:"size:100"`
	EducationLevel *string `gorm:"size:100"`
	AvatarUrl      *string `gorm:"default:null"`
	BackgroundUrl  *string `gorm:"default:null"`
	Country        *string `gorm:"default:null"`
	City           *string `gorm:"default:null"`
	Address        *string `gorm:"default:null"`
	AboutMe        *string `gorm:"default:null"`
	Roles          []Role  `gorm:"many2many:user_roles;"`
}

// Role represents a role in the system that groups a set of permissions.
type Role struct {
	gorm.Model
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

// Permission represents a specific action or resource that can be accessed.
type Permission struct {
	gorm.Model
	Name        string // Example: "write_article"
	Description string // Example: "Permission to write an article"
}

type Device struct {
	gorm.Model
	UserID          uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"`
	DeviceType      *string `gorm:"size:100"`
	OperatingSystem *string `gorm:"size:100"`
	Browser         *string `gorm:"size:100"`
}

type Series struct {
	gorm.Model
	Title         string  `gorm:"size:255"`
	Description   *string `gorm:"type:text"`
	Category      *string `gorm:"size:100"`
	TotalSeasons  int
	TotalEpisodes int
}

type Season struct {
	gorm.Model
	SeriesID     uint `gorm:"index;not null;constraint:OnDelete:CASCADE"`
	SeasonNumber int
	Description  *string `gorm:"type:text"`
}

type Episode struct {
	gorm.Model
	SeasonID      uint `gorm:"index;not null;constraint:OnDelete:CASCADE"`
	EpisodeNumber int
	Title         string  `gorm:"size:255"`
	Description   *string `gorm:"type:text"`
	Duration      int
	VideoID       uint `gorm:"index;not null;constraint:OnDelete:CASCADE"`
}

type Video struct {
	gorm.Model
	Title           string  `gorm:"size:255"`
	Description     *string `gorm:"type:text"`
	Uploader        *string `gorm:"size:255"`
	Duration        int
	Category        *string    `gorm:"size:100"`
	Resolution      *string    `gorm:"size:100"`
	BelongsToSeries *uint      `gorm:"index;constraint:OnDelete:SET NULL"`
	CoverImageUrl   *string    `gorm:"default:null"`
	VideoURLs       []VideoURL `gorm:"foreignKey:VideoID"` // 关联多个视频URL
}

type VideoURL struct {
	gorm.Model
	VideoID uint   `gorm:"index"`    // 关联的视频ID
	URL     string `gorm:"size:255"` // 视频的URL
}

type Tag struct {
	gorm.Model
	TagName string `gorm:"size:100"`
}

type VideoTag struct {
	gorm.Model
	VideoID uint `gorm:"index;not null"`
	TagID   uint `gorm:"index;not null"`
}

type Advertisement struct {
	gorm.Model
	Advertiser string    `gorm:"size:255"`  // 广告商
	Content    string    `gorm:"type:text"` // 广告内容
	StartDate  time.Time // 广告开始时间
	EndDate    time.Time // 广告结束时间
	IsActive   bool      `gorm:"default:false"` // 广告是否激活
}
