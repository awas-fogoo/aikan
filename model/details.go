package model

import "github.com/jinzhu/gorm"

type Details struct {
	gorm.Model
	//所属视频
	Vid uint `gorm:"unique;not null"`
	//分P使用的标题
	Title string `gorm:"type:varchar(50);"`
	//不同分辨率
	P360  string `gorm:"type:varchar(255);"`
	P480  string `gorm:"type:varchar(255);"`
	P720  string `gorm:"type:varchar(255);"`
	P1080 string `gorm:"type:varchar(255);"`
	P2048 string `gorm:"type:varchar(255);"`
	P4096 string `gorm:"type:varchar(255);"`
	//不对分辨率进行处理使用原始分辨率
	Original string  `gorm:"type:varchar(255);"`
	Duration float64 `gorm:"default:0"`      //视频时长
	Review   int     `gorm:"not null;index"` //审核状态
}
