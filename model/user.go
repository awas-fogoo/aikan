package model

import (
	"gorm.io/gorm"
	"time"
)

/*  mysql  */

type User struct {
	gorm.Model
	Username       string  `gorm:"size:255;unique"` //用户名
	Nickname       *string `gorm:"size:255"`        //昵称
	Email          string  `gorm:"size:255;unique"` //邮箱
	Password       string  `gorm:"size:255"`        //密码
	Age            *int    //年龄
	Gender         *string //性别
	Occupation     *string `gorm:"size:100"`              //职业
	EducationLevel *string `gorm:"size:100"`              //教育水平
	AvatarUrl      *string `gorm:"default:null"`          //头像
	BackgroundUrl  *string `gorm:"default:null"`          //背景图片地址
	Country        *string `gorm:"default:null"`          //国家
	City           *string `gorm:"default:null"`          //城市
	Address        *string `gorm:"default:null"`          //地址
	AboutMe        *string `gorm:"default:null"`          //个人介绍
	Roles          []Role  `gorm:"many2many:user_roles;"` //身份
}

// Role represents a role in the system that groups a set of permissions.
type Role struct {
	gorm.Model
	Name        string       //身份名称
	Permissions []Permission `gorm:"many2many:role_permissions;"` // 权限
}

// Permission represents a specific action or resource that can be accessed.
type Permission struct {
	gorm.Model
	Name        string // Example: "write_article"
	Description string // Example: "Permission to write an article"
}

type Device struct {
	gorm.Model
	UserID          uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"` //用户ID
	DeviceType      *string `gorm:"size:100"`                                   //设备类型
	OperatingSystem *string `gorm:"size:100"`                                   //操作系统
	Browser         *string `gorm:"size:100"`                                   //浏览器
}

type Series struct {
	gorm.Model
	Title         string  `gorm:"size:255"`  //标题
	Description   *string `gorm:"type:text"` //描述
	Category      *string `gorm:"size:100"`  //分类
	TotalSeasons  int     //总季数
	TotalEpisodes int     //总集数
}

type Season struct {
	gorm.Model
	SeriesID     uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"` //系列ID
	SeasonNumber int     //季编号
	Description  *string `gorm:"type:text"` //描述
}

type Episode struct {
	gorm.Model
	SeasonID      uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"` //季ID
	EpisodeNumber int     //集编号
	Title         string  `gorm:"size:255"`  //标题
	Description   *string `gorm:"type:text"` //描述
	Duration      int     //时长
	VideoID       uint    `gorm:"index;not null;constraint:OnDelete:CASCADE"` //视频ID
}

type Video struct {
	gorm.Model
	Title            string     `gorm:"size:255"`  //标题
	Description      *string    `gorm:"type:text"` //描述
	Uploader         *string    `gorm:"size:255"`  //上传者
	Duration         int        //时长
	Category         []Category `gorm:"onDelete:CASCADE"`                   //分类
	Resolution       *string    `gorm:"size:100"`                           //分辨率
	BelongsToSeries  *uint      `gorm:"index;constraint:OnDelete:SET NULL"` //属于系列
	CoverImageUrl    *string    `gorm:"default:null"`                       //封面地址
	IsRecommend      int        //是否推荐，可以放在热播，或轮播图 0是不推荐 ，1是推荐 默认是0
	CollectionNumber int        //集数，这个可以后台定时任务来计算集数（也可以在更新集数后直接调用）
	VideoURLs        []VideoURL `gorm:"foreignKey:VideoID"` // 关联多个视频URL
}

type Category struct {
	gorm.Model
	CategoryName string `gorm:"size:100"` // 分类名字
}

type VideoURL struct {
	gorm.Model
	VideoID uint   `gorm:"index"`    // 关联的视频ID
	URL     string `gorm:"size:255"` // 视频的URL
	Name    string
}

type Tag struct {
	gorm.Model
	TagName string `gorm:"size:100"` // 标签
}

type VideoTag struct {
	gorm.Model
	VideoID uint `gorm:"index;not null"` //视频标签
	TagID   uint `gorm:"index;not null"` //标签ID
}

type Advertisement struct {
	gorm.Model
	Advertiser string    `gorm:"size:255"`  // 广告商
	Content    string    `gorm:"type:text"` // 广告内容
	StartDate  time.Time // 广告开始时间
	EndDate    time.Time // 广告结束时间
	IsActive   bool      `gorm:"default:false"` // 广告是否激活
}
