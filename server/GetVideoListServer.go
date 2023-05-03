package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
)

type GetVideoListDTO struct {
	ID          uint    `json:"id"`
	CreatedAt   int64   `json:"created_at"`
	Title       string  `json:"title"`
	Url         string  `json:"url"`
	CoverUrl    string  `json:"cover_url"`
	Views       uint    `json:"views"`
	Likes       uint    `json:"likes"`
	Collections uint    `json:"collections"`
	Duration    float64 `json:"duration"`
	UserID      uint    `json:"user_id"`
	AvatarUrl   string  `json:"avatar_url"`
	Nickname    string  `json:"nickname"`
}

func ToGetVideoListDTOs(vs []model.Video) []GetVideoListDTO {
	dtos := make([]GetVideoListDTO, len(vs))
	for i, v := range vs {
		time := v.CreatedAt.Unix()
		dtos[i] = GetVideoListDTO{ID: v.ID,
			CreatedAt:   time,
			Title:       v.Title,
			Url:         v.Url,
			CoverUrl:    v.CoverUrl,
			Views:       v.Views,
			Likes:       v.Likes,
			Collections: v.Collections,
			Duration:    v.Duration,
			UserID:      v.UserID,
			AvatarUrl:   v.User.AvatarUrl,
			Nickname:    v.User.Nickname}
	}
	return dtos
}
func GetVideoListServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	var videos []model.Video
	db.Preload("User").Order("RAND()").Limit(10).Find(&videos)
	// 还需要修改 todo
	c.JSON(0, dto.Success(ToGetVideoListDTOs(videos)))

}
