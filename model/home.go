package model

import "github.com/jinzhu/gorm"

type SwiperList struct {
	gorm.Model
	Uid       uint   `gorm:"unique;not null" json:"uid"`
	ImgUrl    string `gorm:"varchar(255);not null" json:"imgUrl"`
	VideoHref string `gorm:"varchar(255);not null" json:"videoHref"`
}
