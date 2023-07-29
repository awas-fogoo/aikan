package controller

import (
	"awesomeProject0511/services"
	"github.com/gin-gonic/gin"
)

func VideoController(c *gin.Context) {
	services.UploadVideoServer(c)
}
func GetVideoListController(c *gin.Context) {
	services.GetVideoListServer(c)
}

func GetVideoDetailController(c *gin.Context) {
	services.GetVideoDetailServer(c)
}

func SearchVideoController(c *gin.Context) {
	services.SearchVideoServer(c)
}

func AddCollectionVideoController(c *gin.Context) {
	services.AddCollectionVideoServer(c)
}

func AddLikeVideoController(c *gin.Context) {
	services.AddLikeVideoServer(c)
}

func AddCommentVideoController(c *gin.Context) {
	services.AddCommentVideoServer(c)
}

func GetCommentVideoController(c *gin.Context) {
	services.GerCommentVideoServer(c)
}

func GetHotVideoController(c *gin.Context) {
	services.GetHotVideoServer(c)
}

func VideoAddressForwardController(c *gin.Context) {
	services.VideoStreamServer(c)
}
