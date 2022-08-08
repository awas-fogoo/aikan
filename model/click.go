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
	Loc      uint   `gorm:"default:0"`
	Ip       string `gorm:"type:varchar(50)"`
	Mac      string `gorm:"type:varchar(100)"`
}
