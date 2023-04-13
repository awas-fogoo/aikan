package controller

import (
	"awesomeProject0511/server"
	"github.com/gin-gonic/gin"
)

func VideoController(c *gin.Context) {
	server.UploadVideoServer(c)
}
func GetVideoListController(c *gin.Context) {
	server.GetVideoListServer(c)
}

func GetVideoDetailController(c *gin.Context) {
	server.GetVideoDetailServer(c)
}

func SearchVideoController(c *gin.Context) {
	server.SearchVideoServer(c)
}

func AddCollectionVideoController(c *gin.Context) {
	server.AddCollectionVideoServer(c)
}

func AddLikeVideoController(c *gin.Context) {
	server.AddLikeVideoServer(c)
}

func AddCommentVideoController(c *gin.Context) {
	server.AddCommentVideoServer(c)
}

func GetCommentVideoController(c *gin.Context) {
	server.GerCommentVideoServer(c)
}

func GetHotVideoController(c *gin.Context) {
	server.GetHotVideoServer(c)
}
