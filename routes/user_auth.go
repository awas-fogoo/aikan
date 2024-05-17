package routes

import (
	"github.com/gin-gonic/gin"
	"one/controller"
)

func UserAuthenticationRoute(v1 *gin.RouterGroup) {
	// 发送验证码
	v1.POST("/auth/reg/code", controller.SendVerificationCode)
	v1.POST("/auth/register", controller.Register)
	v1.POST("/auth/login", controller.Login)
	v1.POST("/auth/refresh")
	v1.POST("/auth/logout")

	//获取电影信息列表
	//v1.POST("/auth/getVideo", controller.GetVideo)
}
