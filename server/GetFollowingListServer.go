package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetFollowingListServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	followingID := util.StringToUint(c.Param("user_id"))
	followingList, err := GetFollowingList(db, followingID)
	if err != nil {
		log.Println(err)
		c.JSON(0, dto.Error(-1, "获取关注列表失败，请重试"))
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
