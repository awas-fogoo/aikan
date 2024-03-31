package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"one/vo"
	"testing"
)

//func convertToVideoVo(v *model.Video) vo.SearchVideoVo {
//	var tags []string
//	for _, tag := range v.Tag {
//		tags = append(tags, tag.Name)
//		fmt.Println(tag)
//	}
//
//	return vo.SearchVideoVo{
//		ID:          v.ID,
//		Title:       v.Title,
//		Description: v.Description,
//		Duration:    v.Duration,
//		CreatedAt:   v.CreatedAt,
//		Tags:        "tags",
//	}
//}
func TestSearchVideoServer(t *testing.T) {
	db, err := gorm.Open("mysql", "root:aikan_root_980002_admin@tcp(127.0.0.1:3306)/aikan?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("failed to connect database: %s", err)
	}
	defer db.Close()
	//db.LogMode(true)

	var videos []vo.SearchVideoVo
	q := "1"
	// 分页，默认1-10
	page := 1
	perPage := 10

	// 排序 默认根据创建时间的 降序
	sortField := "created_at"
	sortOrder := "desc"

	// 筛选
	durationMin := 0
	durationMax := 999
	db.Debug().Table("videos").
		Select("videos.*, GROUP_CONCAT(tags.name SEPARATOR ', ') as tags").
		Joins("JOIN video_tags ON videos.id = video_tags.video_id").
		Joins("JOIN tags ON tags.id = video_tags.tag_id").
		Where("videos.deleted_at IS NULL AND duration >= ? AND duration <= ? AND (videos.title LIKE ? OR videos.description LIKE ? OR tags.name LIKE ?)", durationMin, durationMax, "%"+q+"%", "%"+q+"%", "%"+q+"%").
		Group("videos.id").
		Order(sortField + " " + sortOrder).
		Offset((page - 1) * perPage).
		Limit(perPage).
		Scan(&videos)
	//db.Preload("Tags").Where("duration >= ? AND duration <= ? AND (title LIKE ? OR description LIKE ?)", durationMin, durationMax, "%"+q+"%", "%"+q+"%").Order(sortField + " " + sortOrder).Offset((page - 1) * perPage).Limit(perPage).Find(&videos)
	fmt.Println(videos)

	//var videoVos []videoVo
	//for _, v := range videos {
	//	videoVos = append(videoVos, convertToVideoVo(&v))
	//}
}
