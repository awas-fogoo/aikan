package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"github.com/jinzhu/gorm"
)

func (VideoService) AddLikeVideo(userID, videoID uint) *dto.RetDTO {
	db := common.DB

	// 检查视频是否存在
	var video model.Video
	if err := db.Where("id = ?", videoID).First(&video).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.Error(-1, "Video ID does not exist")
		} else {
			return dto.Error(-1, "Failed to fetch video")
		}
	}

	// 查询记录，包括已经软删除的记录
	var userLike model.UserLike
	err := db.Unscoped().Where("user_id = ? AND video_id = ?", userID, videoID).First(&userLike).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.Error(-1, "Failed to check user likes")
	}

	switch {
	case userLike.DeletedAt != nil:
		// 如果存在软删除，则恢复到初始状态
		if err := db.Unscoped().Model(&userLike).Update("deleted_at", gorm.Expr("NULL")).Error; err != nil {
			return dto.Error(-1, "Failed to restart like")
		} else {
			return dto.Success("Restart like success")
		}
	case err == gorm.ErrRecordNotFound:
		// 否则新建记录
		userLike = model.UserLike{
			UserID:  userID,
			VideoID: videoID,
		}
		if err := db.Create(&userLike).Error; err != nil {
			return dto.Error(-1, "Failed to create like")
		} else {
			return dto.Success("Create like success")
		}
	default:
		// 软删除记录
		if err := db.Delete(&userLike).Error; err != nil {
			return dto.Error(-1, "Failed to cancel like")
		} else {
			return dto.Success("Cancel like success")
		}
	}
}
