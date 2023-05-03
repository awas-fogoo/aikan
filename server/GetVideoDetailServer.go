package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
)

type VideoDetailDTO struct {
	ID          uint    `json:"id"`
	CreatedAt   string  `json:"created_at"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Url         string  `json:"url"`
	CoverUrl    string  `json:"cover_url"`
	Views       uint    `json:"views"`
	Likes       uint    `json:"likes"`
	Collections uint    `json:"collections"`
	Duration    float64 `json:"duration"`
	PartitionID uint    `json:"partition_id"`
	Quality     string  `json:"quality"`
	CategoryID  uint    `json:"category_id"`
	UserID      uint    `json:"user_id"`
	Tags        string  `json:"tags"`
	AvatarUrl   string  `json:"avatar_url"`
	Nickname    string  `json:"nickname"`
}

func ToVideoDetailDTO(v model.Video) VideoDetailDTO {
	time := v.CreatedAt.Format("2006-01-02 15:04:05")
	return VideoDetailDTO{
		ID:          v.ID,
		CreatedAt:   time,
		Title:       v.Title,
		Description: v.Description,
		Url:         v.Url,
		CoverUrl:    v.CoverUrl,
		Views:       v.Views,
		Likes:       v.Likes,
		Collections: v.Collections,
		Duration:    v.Duration,
		PartitionID: v.PartitionID,
		Quality:     v.Quality,
		CategoryID:  v.CategoryID,
		UserID:      v.UserID,
		Tags:        "",
		AvatarUrl:   v.User.AvatarUrl,
		Nickname:    v.User.Nickname,
	}
}
func GetVideoDetailServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	id := c.Param("id")
	var video model.Video
	//var videoDetailVos []vo.VideoDetailVo

	// 查询视频信息
	db.Preload("User").Where("id = ?", id).Find(&video)
	//if err := db.Model(&video).Where("id = ?", id).Find(&video).Scan(&videoDetailVos).Error; err != nil {
	//	c.JSON(500, dto.Error(0, "get video detail failed"))
	//	return
	//}
	//
	// 更新视频浏览次数
	video.Views++
	if err := db.Save(&video).Error; err != nil {
		c.JSON(500, dto.Error(0, "update video view failed"))
		return
	}
	c.JSON(0, dto.Success(ToVideoDetailDTO(video)))
}
