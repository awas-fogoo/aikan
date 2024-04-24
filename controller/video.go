package controller

import (
	"github.com/gin-gonic/gin"
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
