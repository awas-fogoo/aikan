package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"github.com/gin-gonic/gin"
	"log"
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
