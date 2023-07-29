package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func AddLikeVideoServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	user, _ := c.Get("user")
	userDto := dto.ToUserDTO(user.(model.User))
	videoID := util.StringToUint(c.Param("id"))
	var video model.Video
	if err := db.Where("id = ?", videoID).First(&video).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(0, dto.Error(-1, "video id does not exist"))
			return
		}
		log.Println(err)
		return
	}
	Like(db, c, userDto.ID, videoID)
}

func Like(db *gorm.DB, c *gin.Context, userID, videoID uint) {
	var userLike model.UserLike

	// 查询记录，包括已经软删除的记录
	err := db.Unscoped().Where("user_id = ? AND video_id = ?", userID, videoID).First(&userLike).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return
	}

	// 如果存在软删除，则恢复到初始状态
	if userLike.DeletedAt != nil {
		// 使用Unscoped方法更新deleted_at字段的值
		if err := db.Unscoped().Model(&userLike).Update("deleted_at", gorm.Expr("NULL")).Error; err != nil {
			log.Println(err)
			return
		}
		c.JSON(0, dto.Success("re like success"))
		return
	} else if err == gorm.ErrRecordNotFound {
		// 否则新建记录
		userLike = model.UserLike{
			UserID:  userID,
			VideoID: videoID,
		}
		if err := db.Create(&userLike).Error; err != nil {
			log.Println(err)
			return
		}
		c.JSON(0, dto.Success("create like success"))
		return
	}

	// 软删除记录
	if err := db.Delete(&userLike).Error; err != nil {
		log.Println(err)
		return
	}
	c.JSON(0, dto.Success("cancel like success"))
	return
}
