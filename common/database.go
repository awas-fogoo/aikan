package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"one/model"
)

var DB *gorm.DB

func InitDB() {
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
		log.Fatalf("Database connection failed")
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatalf(fmt.Sprintf("Database transaction failed: %v", r))
		} else {
			tx.Commit()
		}
	}()

	if err := tx.AutoMigrate(&model.User{}, &model.Video{}, &model.Category{},
		&model.Comment{}, &model.CommentRelation{}, &model.Role{}, &model.Permission{},
		&model.RolePermission{}, &model.Danmuku{},
		&model.UserLike{}, &model.UserCollection{}, &model.Tag{}, &model.VideoTag{},
		&model.Auth{}, &model.SearchRecord{},
	).Error; err != nil {
		log.Fatalf("Unable to migrate table:" + err.Error())
	}
	tx.Model(&model.User{}).AddUniqueIndex("idx_user_username", "username")
	tx.Model(&model.Video{}).AddIndex("idx_video_title", "title")
	tx.Model(&model.Video{}).AddIndex("idx_video_description", "description")
	tx.Model(&model.Video{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	tx.Model(&model.Video{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	tx.Model(&model.Comment{}).AddIndex("idx_comment_user_video", "user_id", "video_id")
	tx.Model(&model.CommentRelation{}).AddIndex("idx_comment_ancestor_descendant", "ancestor_id", "descendant_id")

	db.DB().SetMaxIdleConns(10)  // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(100) // 设置连接池中的最大打开连接数

	DB = db
}
