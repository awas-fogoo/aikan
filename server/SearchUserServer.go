package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
	"log"
)

func SearchUserServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	q := c.Query("q")
	var users []model.User
	var userVos []vo.UserVo
	err := db.Model(&users).Where("username LIKE ?", "%"+q+"%").Or("nickname LIKE ?", "%"+q+"%").Find(&users).Scan(&userVos).Error
	if err != nil {
		log.Panicln("search user failed:", err)
		return
	}
	c.JSON(0, dto.Success(userVos))
}
