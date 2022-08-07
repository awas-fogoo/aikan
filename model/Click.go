package model

import (
	"github.com/jinzhu/gorm"
)

type ChannelClicks struct {
	gorm.Model
	Click    uint   `gorm:"default:0"`
	Uid      string `gorm:"type:varchar(100);not null;unique"`
	Vid      string `gorm:"type:varchar(100);not null"`
	Duration uint   `gorm:"default:0"`
	loc      uint   `gorm:"default:0"`
	ip       string `gorm:"type:varchar(50)"`
	mac      string `gorm:"type:varchar(100)"`
}
