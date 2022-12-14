package common

import (
	"awesomeProject0511/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"net/url"
)

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	var args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect database ,err :" + err.Error())
	}

	// 用户数据库
	db.AutoMigrate(&model.User{})
	// 轮播图数据库
	db.AutoMigrate(&model.SwiperList{})
	{
		// channel
		db.AutoMigrate(&model.ChannelVideo{})
		//db.AutoMigrate(&model.Details{})
		// channel recommend
		db.AutoMigrate(&model.ChannelRecommend{})
		// channel userinfo
		db.AutoMigrate(&model.UserInfo{})
		// channel video info num
		db.AutoMigrate(&model.ChannelVideoInfoNum{})

		//db.AutoMigrate(&model.ChannelLiked{})
		//db.AutoMigrate(&model.ChannelClicks{})
	}

	return db
}

/*

//安装MySQL驱动
go get -u gorm.io/driver/mysql
//安装gorm包
go get -u gorm.io/gorm

*/
