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
	Time        uint      `gorm:"default:0"` // time
}

type ChannelRecommend struct {
	gorm.Model
	Uid      string `gorm:"not null;unique"`
	ImgUrl   string `gorm:"not null;unique"`
	VideoUrl string `gorm:"not null;unique"`
	Title    string `gorm:"not null;"`
}

type ChannelVideoInfoNum struct {
	gorm.Model
	Click      uint   `gorm:"not null;default:0"`
	Danmu      uint   `gorm:"not null;default:0"`
	Like       uint   `gorm:"not null;default:0"`
	Dislike    uint   `gorm:"not null;default:0"`
	Collection uint   `gorm:"not null;default:0"`
	Comment    uint   `gorm:"not null;default:0"`
	Share      uint   `gorm:"not null;default:0"`
	Vid        string `gorm:"not null;unique"`
	Uid        uint   `gorm:"not null;unique"`
}

type ChannelClicks struct {
	gorm.Model
	Click    uint   `gorm:"default:0"`
	Uid      uint   `gorm:"not null;unique"`
	Vid      string `gorm:"type:varchar(100);not null"`
	Duration uint   `gorm:"default:0"`
	Loc      uint   `gorm:"default:0"`
	Ip       string `gorm:"type:varchar(50)"`
	Mac      string `gorm:"type:varchar(100)"`
}

type Danmaku struct {
	gorm.Model
	Vid   uint   `gorm:"not null;index"`
	Part  uint   `gorm:"default:1;index"`
	Time  uint   `gorm:"not null"`  //时间
	Type  int    `gorm:"default:0"` //类型0滚动;1顶部;2底部
	Color string `gorm:"type:varchar(10);default:'#fff'"`
	Text  string `gorm:"type:varchar(100);not null"`
	Uid   uint   `gorm:"not null"`
}

type Details struct {
	gorm.Model
	//所属视频
	Vid uint `gorm:"unique;not null"`
	//分P使用的标题
	Title string `gorm:"type:varchar(50);"`
	//不同分辨率
	P360  string `gorm:"type:varchar(255);"`
	P480  string `gorm:"type:varchar(255);"`
	P720  string `gorm:"type:varchar(255);"`
	P1080 string `gorm:"type:varchar(255);"`
	P2048 string `gorm:"type:varchar(255);"`
	P4096 string `gorm:"type:varchar(255);"`
	//不对分辨率进行处理使用原始分辨率
	Original string  `gorm:"type:varchar(255);"`
	Duration float64 `gorm:"default:0"`      //视频时长
	Review   int     `gorm:"not null;index"` //审核状态
}

type ChannelLiked struct {
	gorm.Model
	Uid    uint   `gorm:"not null"`
	Vid    string `gorm:"not null"`
	Status bool   `gorm:"default:false"` //是否点赞
}

type Partition struct {
	gorm.Model
	Content  string `gorm:"varchar(20);"`
	ParentId uint   `gorm:"default:0"` //所属分区ID
}
