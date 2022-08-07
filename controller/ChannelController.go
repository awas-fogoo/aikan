package controller

import (
	"awesomeProject0511/server"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetTopChannelDetail(c *gin.Context) {
	videoUrl := c.Param("id")
	res := server.GetChannelService(videoUrl)
	c.JSON(200, res)
}
