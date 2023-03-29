package model

import "github.com/jinzhu/gorm"

type VideoRelation struct {
	gorm.Model
	Uid       uint   `gorm:"not null"`
	IsLike    bool   `gorm:"not null"`
	IsDislike bool   `gorm:"not null"`
	IsCollect bool   `gorm:"not null"`
	IsFollow  bool   `gorm:"not null"`
	IsShare   bool   `gorm:"not null"`
	Vid       string `gorm:"not null"`
}
