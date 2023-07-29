package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
)

type VideoService struct{}

type VideoBrief struct {
	ID       uint   `json:"id"`
	CoverURL string `json:"cover_url"`
	Views    uint   `json:"views"`
}

func (s VideoService) GetCollections(userID uint) ([]VideoBrief, error) {
	// 连接数据库
	db := common.InitDB()
	defer db.Close()

	// 查询收藏列表
	var videos []model.Video
	var videoBriefs []VideoBrief
	if err := db.
		Joins("INNER JOIN user_collections AS uc ON videos.id = uc.video_id").
		Where("uc.user_id = ?", userID).
		Order("uc.created_at DESC").
		Limit(10).
		Select("videos.id, videos.cover_url, videos.views").
		Find(&videos).
		Scan(&videoBriefs).Error; err != nil {
		return nil, err
	}

	return videoBriefs, nil
}

func (s VideoService) GetLikes(userID uint) ([]VideoBrief, error) {
	// 连接数据库
	db := common.InitDB()
	defer db.Close()

	// 查询收藏列表
	var videos []model.Video
	var videoBriefs []VideoBrief
	if err := db.
		Joins("INNER JOIN user_likes AS uc ON videos.id = uc.video_id").
		Where("uc.user_id = ?", userID).
		Order("uc.created_at DESC").
		Limit(10).
		Select("videos.id, videos.cover_url, videos.views").
		Find(&videos).
		Scan(&videoBriefs).Error; err != nil {
		return nil, err
	}

	return videoBriefs, nil
}
