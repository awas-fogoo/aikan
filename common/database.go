package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
	"one/model"
)

var DB *gorm.DB

func InitDB() {
	// 从配置文件读取数据库连接信息
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	// 构建 DSN
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	// 使用 mysql.New 函数构建 MySQL Dialector
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Successfully connected to Mysql")
	//自动迁移所有模型
	models := []interface{}{&model.User{}, &model.Role{}, &model.Permission{},
		&model.Device{}, &model.Series{}, &model.Season{}, &model.Episode{},
		&model.Video{}, &model.Tag{}, &model.VideoTag{}, &model.VideoURL{},
		&model.Advertisement{}, &model.Story{}}
	for _, _model := range models {
		err = db.AutoMigrate(_model)
		if err != nil {
			fmt.Println("Failed to migrate database model:", err)
			return
		}
	}
	fmt.Println("Database migration completed successfully.")
	DB = db
}
