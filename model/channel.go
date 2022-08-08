package model

import "github.com/jinzhu/gorm"

type ChannelList struct {
	gorm.Model
	ImgUrl     string `gorm:"varchar(255);not null" json:"imgUrl"`
	Hot        string `gorm:"varchar(20)" json:"hot"`
	Title      string `gorm:"varchar(100);not null" json:"title"`
	HeadUrl    string `gorm:"varchar(255);not null" json:"headUrl"`
	Name       string `gorm:"varchar(100);not null" json:"name"`
	Fans       uint   `gorm:"int(100)" json:"fans"`
	Follow     bool   `gorm:"default:0" json:"follow"`
	See        uint   `gorm:"int(255);not null;default:1" json:"see"`
	Danmu      uint   `gorm:"int(200)" json:"danmu"`
	VideoUrl   string `gorm:"varchar(255);not null;unique" json:"videoUrl"`
	Desc       string `gorm:"varchar(255);not null" json:"desc"`
	Like       uint   `gorm:"int(255)" json:"like"`
	Dislike    uint   `gorm:"int(100)" json:"dislike"`
	Collection uint   `gorm:"int(200)" json:"collection"`
	Share      uint   `gorm:"int(50);" json:"share"`
	Comment    uint   `gorm:"int(255)" json:"comment"`
	Uid        string `gorm:"varchar(100);not null" json:"uid"`
}
