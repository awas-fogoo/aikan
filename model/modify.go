package model

import "github.com/jinzhu/gorm"

type Archive struct {
	gorm.Model
	Follow   string `gorm:"varchar(300)";json:"follow"`
	History  string `gorm:"varchar(1000)";json:"history"`
	Like     string `gorm:"varchar(300)";json:"like"`
	Dislike  string `gorm:"varchar(100)";json:"dislike"`
	Favorite string `gorm:"varchar(300)";json:"favorite"`
	Time     Time   `json:"time"`
	Uid      string `gorm:"varchar(50)";json:"uid"`
}
