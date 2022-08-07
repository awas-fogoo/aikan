package model

import "github.com/jinzhu/gorm"

type ChannelList struct {
	gorm.Model
	Cover       string    `gorm:"varchar(255);not null"`
	Hot         string    `gorm:"varchar(20)" json:"hot"`
	Title       string    `gorm:"varchar(100);not null"`
	Videos      []Details `gorm:"ForeignKey:vid;AssociationForeignKey:vid"`
	Desc        string    `gorm:"varchar(255);default:'什么都没有~^v^~'"`
	Vid         string    `gorm:"varchar(100);unique;not null"`
	Uid         string    `gorm:"varchar(100);unique;not null"`
	Copyright   bool      `gorm:"not null"`  //是否为原创(版权)
	Weights     float32   `gorm:"default:0"` //视频权重(目前还没使用)
	Clicks      uint      `gorm:"default:0"` //点击量
	Review      int       `gorm:"not null"`  //审核状态
	PartitionID uint      `gorm:"default:0"` //分区ID
}
