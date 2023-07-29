package routes

import (
	"awesomeProject0511/controller"
	"awesomeProject0511/middleware"
	"github.com/gin-gonic/gin"
)

func VideoRoute(v1 *gin.RouterGroup) {
	v1.GET("/channel", controller.VideoAddressForwardController)
	v1.POST("/videos", middleware.AuthMiddleware(), controller.VideoController)
	v1.GET("/videos", controller.GetVideoListController)
	v1.GET("/video/:id", controller.GetVideoDetailController)
	v1.GET("/videos/search", controller.SearchVideoController)
	v1.POST("/video/:id/like", middleware.AuthMiddleware(), controller.AddLikeVideoController)
	v1.POST("/video/:id/collection", middleware.AuthMiddleware(), controller.AddCollectionVideoController)
	v1.POST("/video/comment", middleware.AuthMiddleware(), controller.AddCommentVideoController)
	v1.GET("/video/:id/comments", controller.VideosController{}.GetComments)
	v1.POST("/video/danmu/:vid", controller.VideosController{}.AddDanmu)
	v1.GET("/video/danmu/:vid", controller.VideosController{}.GetDanmu)
	v1.GET("/videos/hot", controller.GetHotVideoController)
	v1.GET("/videos/:id/collections", controller.VideosController{}.GetCollections)
	v1.GET("/videos/:id/likes", controller.VideosController{}.GetLikes)
}

/*
在RESTful API中，通常使用以下命名约定：

GET /resource：获取资源列表，常常命名为 Index
GET /resource/:id：获取单个资源，常常命名为 Show
POST /resource：创建新资源，常常命名为 Create
PUT /resource/:id：更新单个资源，常常命名为 Update
DELETE /resource/:id：删除单个资源，常常命名为 Delete

*/
