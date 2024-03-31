package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"one/common"
	"one/dto"
	"one/util"
)

func GetFollowingListServer(c *gin.Context) {
	db := common.DB
	followingID := util.StringToUint(c.Param("user_id"))
	followingList, err := GetFollowingList(db, followingID)
	if err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(5000, "获取关注列表失败，请重试"))
	}
	var usersJSON []map[string]interface{}
	for _, following := range followingList {
		userJSON := map[string]interface{}{
			"nickname":   following.Nickname,
			"avatar_url": following.AvatarUrl,
			"about_me":   following.AboutMe,
		}
		usersJSON = append(usersJSON, userJSON)
	}

	c.JSON(200, dto.Success(usersJSON))
}
