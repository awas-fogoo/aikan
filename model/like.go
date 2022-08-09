package model

import "github.com/jinzhu/gorm"

type ChannelLiked struct {
	gorm.Model
	Uid    uint   `gorm:"not null"`
	Vid    string `gorm:"not null"`
	Status bool   `gorm:"default:false"` //是否点赞
}
