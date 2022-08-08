package model

type SwiperList struct {
	Id        string `gorm:"primary_key" json:"id"`
	Uid       string `gorm:"int(100);unique;not null" json:"uid"`
	ImgUrl    string `gorm:"varchar(255);not null" json:"imgUrl"`
	VideoHref string `gorm:"varchar(255);not null" json:"videoHref"`
}
