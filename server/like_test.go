package server

import (
	"awesomeProject0511/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	"time"
)

func TestLike(t *testing.T) {
	db, err := gorm.Open("mysql", "root:aikan_root_980002_admin@tcp(127.0.0.1:3306)/aikan?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("failed to connect database: %s", err)
	}
	defer db.Close()

	// 清空测试数据
	db.Delete(&model.UserLike{})
	db.Delete(&model.UserCollection{})

	// 创建测试数据
	userID := uint(14)
	videoID := uint(3)
	var userLike model.UserLike
	err = db.Where("user_id = ? AND video_id = ?", userID, videoID).First(&userLike).Error
	fmt.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return
	}

	if userLike.DeletedAt == nil {
		userLike.DeletedAt = &time.Time{} // 初始化 DeletedAt 字段
		if err := db.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&userLike).Error; err != nil {
			return
		}
		fmt.Println("删除点赞")
	} else {
		fmt.Println("点赞已取消")
	}
	if userLike.DeletedAt != nil {
		if err := db.Unscoped().Model(&userLike).Where("user_id = ? AND video_id = ?", userID, videoID).Update("deleted_at", gorm.Expr("NULL")).Error; err != nil {
			return
		}
		fmt.Println("取消点赞成功后再次点赞成功")
		return
	} else if err == gorm.ErrRecordNotFound {
		// 否则新建记录
		userLike = model.UserLike{
			UserID:  userID,
			VideoID: videoID,
		}
		fmt.Println("创建点赞成功")
		if err := db.Create(&userLike).Error; err != nil {
			log.Panicln(err, "create")
			return
		}
		log.Println(err)
		return
	}

	//var userLike model.UserLike
	//err = db.Unscoped().Where("user_id = ? AND video_id = ?", userID, videoID).First(&userLike).Error
	//fmt.Println(err)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	fmt.Println(err)
	//}
	////
	//// 如果存在软删除，则恢复到初始状态
	//if userLike.DeletedAt != nil {
	//	if err := db.Unscoped().Model(&userLike).Update("deleted_at", gorm.Expr("NULL")).Error; err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("取消点赞成功后再次点赞成功")
	//	fmt.Println(err)
	//	//return
	//} else if err == gorm.ErrRecordNotFound {
	//	// 否则新建记录
	//	userLike = model.UserLike{
	//		UserID:  userID,
	//		VideoID: videoID,
	//	}
	//	fmt.Println("创建点赞成功")
	//	if err := db.Create(&userLike).Error; err != nil {
	//		fmt.Println(err)
	//	}
	//	log.Println(err)
	//	return
	//}
	//fmt.Println("删除点赞")
	//if err := db.Delete(&userLike).Error; err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(err)

	// 第一次点赞
}
