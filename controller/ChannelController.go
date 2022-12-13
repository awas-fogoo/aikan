package controller

import (
	"awesomeProject0511/server"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetTopChannelDetail(c *gin.Context) {
	videoUrl := c.Param("id")
	/*
		1、对videoUrl进行只能a-zA-z_
		2、获取本机外网ip，
		3、判断是否登入，直接看头部是否有token
		if token 「
			使用token 提取uid（not null unique） 把vid（not null） 查看时间，持续时间，看到那，获取外网ip，存入redis
			如果性能良好的情况就立刻提交，否则随机生成 定时提交
		」
		无token需等待10钟才开始算redis 定时提交

	*/
	// 视频播放页面
	res := server.GetChannelService(videoUrl)
	c.JSON(200, res)
	// 视频播放页面下的recommend

}
