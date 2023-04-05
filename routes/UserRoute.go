package routes

import (
	"awesomeProject0511/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(v1 *gin.RouterGroup) {
	v1.GET("/users/search", controller.SearchUserController)
}
