package services

import (
	"awesomeProject0511/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

//type VideoDetailDTO struct {
//	ID          uint    `json:"id"`
//	CreatedAt   string  `json:"created_at"`
//	Title       string  `json:"title"`
//	Description string  `json:"description"`
//	Url         string  `json:"ulr"`
//	CoverUrl    string  `json:"cover_url"`
//	Views       uint    `json:"views"`
//	Likes       uint    `json:"likes"`
//	Collections uint    `json:"collections"`
//	Duration    float64 `json:"duration"`
//	PartitionID uint    `json:"partition_id"`
//	Quality     string  `json:"quality"`
//	CategoryID  uint    `json:"category_id"`
//	UserID      uint    `json:"user_id"`
//	Tags        string  `json:"tags"`
//	AvatarUrl   string  `json:"avatar_url"`
//	Nickname    string  `json:"nickname"`
//}
//
//func ToVideoDetailDTO(v model.Video) VideoDetailDTO {
//	time := v.CreatedAt.Format("2006-01-02 15:04:05")
//	return VideoDetailDTO{
//		ID:          v.ID,
//		CreatedAt:   time,
//		Title:       v.Title,
//		Description: v.Description,
//		Url:         v.Url,
//		CoverUrl:    v.CoverUrl,
//		Views:       v.Views,
//		Likes:       v.Likes,
//		Collections: v.Collections,
//		Duration:    v.Duration,
//		PartitionID: v.PartitionID,
//		Quality:     v.Quality,
//		CategoryID:  v.CategoryID,
//		UserID:      v.UserID,
//		Tags:        "",
//		AvatarUrl:   v.User.AvatarUrl,
//		Nickname:    v.User.Nickname,
//	}
//}

func TestGetVideoDetailServer(t *testing.T) {
	db, err := gorm.Open("mysql", "root:aikan_root_980002_admin@tcp(127.0.0.1:3306)/aikan?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("failed to connect database: %s", err)
	}
	defer db.Close()
	db.LogMode(true)
	id := 27
	var video model.Video
	//var videoDetailVos []vo.VideoDetailVo
	db.Preload("User").Where("id = ?", id).Find(&video)
	fmt.Println(video)
	fmt.Println(ToVideoDetailDTO(video))
}
