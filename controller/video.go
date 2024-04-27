package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"one/common"
	"one/dto"
	"one/model"
)

func UploadVideo(c *gin.Context) {
	var video model.Video
	var urls []string
	var tags []string

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(200, dto.Error(4001, "error 1"))
		return
	}
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(200, dto.Error(4002, "error 1"))
		return
	}
	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(200, dto.Error(4003, "error 1"))
		return
	}
	db := common.DB

	db.Create(&video)

	for _, url := range urls {
		db.Create(&model.VideoURL{VideoID: video.ID, URL: url})
	}
	for _, tagName := range tags {
		var tag model.Tag
		db.Where(model.Tag{TagName: tagName}).FirstOrCreate(&tag)
		db.Create(&model.VideoTag{VideoID: video.ID, TagID: tag.ID})
	}
	c.JSON(200, dto.Success("Video, URLs, and tags uploaded successfully"))

}
func UploadSeason(c *gin.Context) {
	var (
		seasonData model.Season
		episodes   []model.Episode
	)
	if err := c.ShouldBindJSON(&seasonData); err != nil {
		c.JSON(200, dto.Error(4001, "error 1"))
		return
	}
	if err := c.ShouldBindJSON(&episodes); err != nil {
		c.JSON(200, dto.Error(4002, "error 1"))
		return
	}
	db := common.DB

	db.Create(&seasonData)
	for _, episode := range episodes {
		episode.SeasonID = seasonData.ID
		db.Create(&episode)
	}
	c.JSON(200, dto.Success("Season and episodes uploaded successfully"))
}

// 获取视频
func GetVideo(c *gin.Context) {
	var db *gorm.DB

	var video model.Video
	storyId := c.Param("storyId")

	//db.Where("story_id=?", storyId).First(&video)
	//c.JSON(200, dto.RetDTO{Message: "Login successful", Data: video})
	result := db.Table("videos").Where("story_id = ?", storyId).First(&video)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Video found", "data": video})
}
