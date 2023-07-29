package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func GetFollowersListServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	followersID := util.StringToUint(c.Param("user_id"))
	followersList, err := GetFollowerList(db, followersID)
	if err != nil {
		log.Println(err)
		c.JSON(0, dto.Error(-1, "获取粉丝列表失败，请重试"))
	}
	var usersJSON []map[string]interface{}
	for _, following := range followersList {
		userJSON := map[string]interface{}{
			"nickname":   following.Nickname,
			"avatar_url": following.AvatarUrl,
			"about_me":   following.AboutMe,
		}
		usersJSON = append(usersJSON, userJSON)
	}

	jsonBytes, err := json.Marshal(usersJSON)
	if err != nil {
		log.Println(err)
		c.JSON(500, dto.Error(-1, "服务器内部错误，请重试"))
	}

	c.JSON(200, dto.Success(string(jsonBytes)))
}
