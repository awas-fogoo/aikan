package model

import "github.com/jinzhu/gorm"

type ChannelLiked struct {
	gorm.Model
	Uid    uint    `gorm:"not null"`
	Vid    uint    `gorm:"not null"`
	Status bool    `gorm:"default:false"` //是否点赞
	Video  Details `gorm:"ForeignKey:vid;AssociationForeignKey:vid"`
}
