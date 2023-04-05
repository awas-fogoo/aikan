package routes

import (
	"awesomeProject0511/controller"
	"awesomeProject0511/middleware"
	"github.com/gin-gonic/gin"
)

func VideoRoute(v1 *gin.RouterGroup) {
	v1.POST("/videos", middleware.AuthMiddleware(), controller.VideoController)
	v1.GET("/videos", controller.GetVideoListController)
	v1.GET("/video/:id", controller.GetVideoDetailController)
	v1.GET("/videos/search", controller.SearchVideoController)
	v1.POST("/video/:id/like", middleware.AuthMiddleware(), controller.AddLikeVideoController)
	v1.POST("/video/:id/collection", middleware.AuthMiddleware(), controller.AddCollectionVideoController)
	v1.POST("/video/comment", middleware.AuthMiddleware(), controller.AddCommentVideoController)
	v1.GET("/video/comments", controller.GetCommentVideoController)
}
