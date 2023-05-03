package server

import (
	"awesomeProject0511/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

func TestGetProfileServer(t *testing.T) {
	db, err := gorm.Open("mysql", "root:aikan_root_980002_admin@tcp(127.0.0.1:3306)/aikan?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("failed to connect database: %s", err)
	}
	defer db.Close()
	userID := uint(21)
	userID2 := uint(22)

	user := &model.User{}
	err = db.Preload("Videos").Preload("Following").Preload("Followers").First(user, userID).Error
	if err != nil {
		fmt.Println(err)
	}

	// 获取用户上传视频总数
	count := len(user.Videos)
	fmt.Println(count)

	// 获取用户关注数量
	followingCount := len(user.Following)
	fmt.Println(followingCount)

	// 获取用户粉丝数量
	followerCount := len(user.Followers)
	fmt.Println(followerCount)
	// 查询当前用户是否已关注目标用户

	isFollowing := false
	for _, followee := range user.Following {
		if followee.ID == userID2 {
			isFollowing = true
			break
		}
	}
	fmt.Println(isFollowing)

	// 按照最新上传的时间排序
	videos := []model.Video{}
	db.Where("user_id = ?", userID).Order("created_at DESC").Limit(10).Find(&videos)

	// 将每个视频的 views 和 cover_url 添加到数组中
	results := make([]map[string]interface{}, 0, len(videos))
	for _, video := range videos {
		result := make(map[string]interface{})
		result["id"] = video.ID
		result["views"] = video.Views
		result["cover_url"] = video.CoverUrl
		results = append(results, result)
	}

	// 输出数组结果
	fmt.Println(results)

}
