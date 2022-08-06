package model

import (
	"github.com/jinzhu/gorm"
)

// SwiperList 轮播图
type SwiperList struct {
	gorm.Model
	ImgUrl  string `gorm:"varchar(255);not null"`
	VideoId string `gorm:"varchar(255);unique;not null"`
	Uid     string `gorm:"int(100);unique;not null"`
}
