package controller

import (
	"awesomeProject0511/server"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func HomeSwiperList(c *gin.Context) {
	res := server.GetSwiperListService()
	c.JSON(200, res)
}
