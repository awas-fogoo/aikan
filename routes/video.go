package routes

import (
	"github.com/gin-gonic/gin"
	"one/controller"
)

func VideoRoute(v1 *gin.RouterGroup) {
	//v1.GET("/channel", controller.VideosController{}.VideoAddressForward)
	//v1.POST("/videos", middleware.AuthMiddleware(), controller.VideosController{}.UploadVideo)
	//v1.GET("/videos", controller.VideosController{}.GetVideoList)
	//v1.GET("/video/:id", controller.VideosController{}.GetVideoDetail)
	//v1.GET("/videos/search", controller.VideosController{}.SearchVideo)
	//v1.POST("/video/:id/like", middleware.AuthMiddleware(), controller.VideosController{}.AddLikeVideo)
	//v1.POST("/video/:id/collection", middleware.AuthMiddleware(), controller.VideosController{}.AddCollectionVideo)
	//v1.POST("/video/comment", middleware.AuthMiddleware(), controller.VideosController{}.AddCommentVideo)
	//v1.GET("/video/:id/comments", controller.VideosController{}.GetComments)
	//v1.POST("/video/danmu/:vid", controller.VideosController{}.AddDanmu)
	//v1.GET("/video/danmu/:vid", controller.VideosController{}.GetDanmus)
	//v1.GET("/videos/hot", controller.VideosController{}.GetHotVideo)
	//v1.GET("/videos/:id/collections", controller.VideosController{}.GetCollections)
	//v1.GET("/videos/:id/likes", controller.VideosController{}.GetLikes)
	v1.GET("/video/upload", controller.UploadVideo)
	v1.GET("/video/season", controller.UploadSeason)
	v1.POST("/video/getVideoMsgByVideoId", controller.GetVideoMsgByVideoId)
	v1.POST("/video/getVideoStory", controller.GetVideoStory) //相当与频道
	//v1.POST("/video/getVideoStory", controller.GetVideoStory) //相当与频道
	v1.POST("/video/GetVideoByStoryId", controller.GetVideoByStoryId)
	v1.POST("/video/VideoSearch", controller.VideoSearch)
	v1.POST("/video/GetVideoAllList", controller.GetVideoAllList)
}

/*
在RESTful API中，通常使用以下命名约定：

GET /resource：获取资源列表，常常命名为 Index
GET /resource/:id：获取单个资源，常常命名为 Show
POST /resource：创建新资源，常常命名为 Create
PUT /resource/:id：更新单个资源，常常命名为 Update
DELETE /resource/:id：删除单个资源，常常命名为 Delete

*/
