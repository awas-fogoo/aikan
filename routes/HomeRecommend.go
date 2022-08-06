package routes

import (
	"awesomeProject0511/controller"
	"awesomeProject0511/middleware"
	"github.com/gin-gonic/gin"
)

func HomeRecommend(v1 *gin.RouterGroup) {

	// 获取首页推荐图片列表
	v1.GET("/home", controller.GetHomeList)

	// 获取推荐图片详情播放页面
	v1.GET("/channel/:id", controller.GetTopChannelDetail)

	// 暂时不开发。
	v1.GET("/channel", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"static": "channel_list",
		})
	})

	// 发送验证码
	v1.POST("/auth/reg/code", controller.SendVerificationCode)
	// 用户注册页面
	v1.POST("/auth/register", controller.Register)
	v1.POST("/auth/login", controller.Login)
	v1.GET("/auth/info", middleware.AuthMiddleware(), controller.Info)

	// MY SPACE
	v1.GET("/:id", controller.Myspace)

}
