package routes

import (
	"awesomeProject0511/controller"
	"awesomeProject0511/middleware"
	"github.com/gin-gonic/gin"
)

func ChannelRouter(v1 *gin.RouterGroup) {
	v1.POST("/auth/channel/like", middleware.AuthMiddleware(), controller.Like)

}
