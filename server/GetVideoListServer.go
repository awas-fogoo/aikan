package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func GetVideoListServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	var videos []model.Video
	var videoHomeVos []vo.VideoHomeVo
	db.Model(&videos).Order("created_at desc").Limit(10).Find(&videos).Scan(&videoHomeVos)
	c.JSON(0, dto.Success(videoHomeVos))
}
