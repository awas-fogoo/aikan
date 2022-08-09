package model

import "github.com/jinzhu/gorm"

type ChannelVideo struct {
	gorm.Model
	Cover       string    `gorm:"size:255;not null"`
	Hot         string    `gorm:"varchar(20)" json:"hot"`
	Title       string    `gorm:"type:varchar(50);not null;index"`
	Video       []Details `gorm:"ForeignKey:vid;AssociationForeignKey:vid"`
	Desc        string    `gorm:"type:varchar(200);default:'什么都没有~^v^~'"` //视频简介
	Vid         string    `gorm:"not null;unique"`
	Uid         uint      `gorm:"not null;unique"`
	Copyright   bool      `gorm:"not null"`  //是否为原创(版权)
	Weights     float32   `gorm:"default:0"` //视频权重
	Review      int       `gorm:"not null"`  //审核状态
	PartitionID uint      `gorm:"default:0"` //分区ID
}
