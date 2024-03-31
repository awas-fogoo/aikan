package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"one/common"
	"one/dto"
	"one/model"
	"one/util"
)

func UnFollowUserServer(c *gin.Context) {
	db := common.DB
	user, _ := c.Get("user")
	userDto := dto.ToUserDTO(user.(model.User))
	followingID := util.StringToUint(c.Param("user_id"))
	if err := UnfollowUser(db, userDto.ID, followingID); err != nil {
		log.Println(err)
		c.JSON(200, dto.Error(5000, "取消关注失败，请重试"))
	}
	c.JSON(200, dto.Success("取消关注成功"))
}
