package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30);not null"`
	Email    string `gorm:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
	Uid      uint   `gorm:"not null;unique"`
}

type UserInfo struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null"`
	Fans     uint   `gorm:"default:0"`
	Follow   uint   `gorm:"default:0"`
	HeadImg  string `gorm:"type:varchar(30);not null"`
	Personal string `gorm:"varchar(110);not null;unique"`
	Uid      uint   `gorm:"not null;unique"`
}
