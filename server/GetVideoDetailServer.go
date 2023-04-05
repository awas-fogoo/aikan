package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func GetVideoDetailServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	id := c.Param("id")
	var video model.Video
	var videoDetailVos []vo.VideoDetailVo

	// 查询视频信息
	if err := db.Model(&video).Where("id = ?", id).Find(&video).Scan(&videoDetailVos).Error; err != nil {
		c.JSON(500, dto.Error(0, "get video detail failed"))
		return
	}

	// 更新视频浏览次数
	video.Views++
	if err := db.Save(&video).Error; err != nil {
		c.JSON(500, dto.Error(0, "update video view failed"))
		return
	}

	c.JSON(0, dto.Success(videoDetailVos))
}
