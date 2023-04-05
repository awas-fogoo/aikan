package routes

import (
	"awesomeProject0511/controller"
	"github.com/gin-gonic/gin"
)

func UserAuthenticationRoute(v1 *gin.RouterGroup) {
	// 发送验证码
	v1.POST("/auth/reg/code", controller.SendVerificationCode)
	v1.POST("/auth/register", controller.Register)
	v1.POST("/auth/login", controller.Login)
	v1.POST("/auth/refresh")
	v1.POST("/auth/logout")
}
