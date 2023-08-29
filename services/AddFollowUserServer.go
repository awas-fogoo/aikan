package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func AddFollowUserServer(c *gin.Context) {
	db := common.DB
	user, _ := c.Get("user")
	userDto := dto.ToUserDTO(user.(model.User))
	followingID := util.StringToUint(c.Param("user_id"))
	if userDto.ID == followingID {
		c.JSON(200, dto.Error(4000, "cannot follow oneself"))
		return
	}

	if err := FollowUser(db, userDto.ID, followingID); err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(5000, "关注失败，请重试"))
		return
	}
	c.JSON(200, dto.Success("关注成功"))
}

// FollowUser 添加关注
func FollowUser(db *gorm.DB, userID, followingID uint) error {
	var user, following model.User

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, userID).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&following, followingID).Error; err != nil {
			return err
		}
		if err := tx.Model(&user).Association("Following").Append(&following); err != nil {
			return nil
		}
		if err := tx.Model(&following).Association("Followers").Append(&user); err != nil {
			return nil
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GetFollowingList 查询关注列表
func GetFollowingList(db *gorm.DB, userID uint) ([]*model.User, error) {
	var user model.User
	if err := db.Preload("Following").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Following, nil
}

// GetFollowerList 查询粉丝列表
func GetFollowerList(db *gorm.DB, userID uint) ([]*model.User, error) {
	var user model.User
	if err := db.Preload("Followers").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Followers, nil
}

// CountFollowersAndFollowing 统计关注数和粉丝数
func CountFollowersAndFollowing(db *gorm.DB, userID uint) (int64, int64, error) {
	var followingCount, followerCount int64
	if err := db.Model(&model.User{}).Where("id = ?", userID).Joins("JOIN user_followings ON user_followings.user_id = users.id").Count(&followingCount).Error; err != nil {
		return 0, 0, err
	}
	if err := db.Model(&model.User{}).Where("id = ?", userID).Joins("JOIN user_followers ON user_followers.user_id = users.id").Count(&followerCount).Error; err != nil {
		return 0, 0, err
	}
	return followingCount, followerCount, nil
}

// UnfollowUser 取消关注
func UnfollowUser(db *gorm.DB, followerID, followingID uint) error {
	if followerID == followingID {
		return errors.New("cannot follow oneself")
	}
	var follower, following model.User

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&follower, followerID).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&following, followingID).Error; err != nil {
			return err
		}
		if err := tx.Model(&follower).Association("Following").Delete(&following); err != nil {
			log.Println("Following failed")
			return nil
		}
		if err := tx.Model(&following).Association("Followers").Delete(&follower); err != nil {
			log.Println("Followers failed")
			return nil
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
