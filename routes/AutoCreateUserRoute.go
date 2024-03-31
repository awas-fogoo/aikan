package routes

import (
	"github.com/gin-gonic/gin"
	"one/controller"
)

func AutoCreateUserRoute(v1 *gin.RouterGroup) {
	v1.GET("/auto", controller.AutoCreateUser)
}
