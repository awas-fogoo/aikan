package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
	Uid      string `gorm:"varchar(50);not null"`
}
type RegUser struct {
	Name     string
	Email    string
	Password string
	Code     string
}
