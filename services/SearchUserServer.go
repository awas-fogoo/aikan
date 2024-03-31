package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"one/common"
	"one/dto"
	"one/model"
	"one/vo"
)

func SearchUserServer(c *gin.Context) {
	db := common.DB
	q := c.Query("q")
	var users []model.User
	var userVos []vo.UserVo
	err := db.Model(&users).Where("username LIKE ?", "%"+q+"%").Or("nickname LIKE ?", "%"+q+"%").Find(&users).Scan(&userVos).Error
	if err != nil {
		log.Panicln("search user failed:", err)
		return
	}
	c.JSON(200, dto.Success(userVos))
}
