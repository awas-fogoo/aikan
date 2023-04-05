package routes

import (
	"awesomeProject0511/controller"
	"github.com/gin-gonic/gin"
)

func AutoCreateUserRoute(v1 *gin.RouterGroup) {
	v1.GET("/auto", controller.AutoCreateUser)
}
