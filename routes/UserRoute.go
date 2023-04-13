package routes

import (
	"awesomeProject0511/controller"
	"awesomeProject0511/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(v1 *gin.RouterGroup) {
	v1.GET("/users/search", controller.SearchUserController)
	v1.GET("/users/:user_id/following", middleware.AuthMiddleware(), controller.GetFollowingListController)
	v1.GET("/users/:user_id/followers", middleware.AuthMiddleware(), controller.GetFollowersListController)
	v1.POST("/users/:user_id/follow", middleware.AuthMiddleware(), controller.AddFollowUserController)
	v1.DELETE("/users/:user_id/follow", middleware.AuthMiddleware(), controller.UnFollowUserController)
}
