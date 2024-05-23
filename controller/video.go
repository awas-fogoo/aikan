package controller

import (
	"github.com/gin-gonic/gin"
	"one/common"
	"one/dto"
	"one/model"
	"one/services"
	_ "one/services"
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

//func UploadSeason(c *gin.Context) {
//	var (
//		seasonData model.Season
//		episodes   []model.Episode
//	)
//	if err := c.ShouldBindJSON(&seasonData); err != nil {
//		c.JSON(200, dto.Error(4001, "error 1"))
//		return
//	}
//	if err := c.ShouldBindJSON(&episodes); err != nil {
//		c.JSON(200, dto.Error(4002, "error 1"))
//		return
//	}
//	db := common.DB
//
//	db.Create(&seasonData)
//	for _, episode := range episodes {
//		episode.SeasonID = seasonData.ID
//		db.Create(&episode)
//	}
//	c.JSON(200, dto.Success("Season and episodes uploaded successfully"))
//}

//var db *gorm.DB

// 根据storyId和推荐获取视频固定数量
func GetVideoByStoryId(c *gin.Context) {
	videoList, err := services.GetRecommendVideoList(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video found", "data": videoList})
}

// 根据视频id查询详细信息
func GetVideoMsgByVideoId(c *gin.Context) {
	dateils, err := services.GetVideoMsgByVideoId(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video found", "data": dateils})
}

// 获取视频类型(1.电影 2.电视剧 3.综艺)
//func GetVideoStory(c *gin.Context) {
//	db := common.DB
//	var Story []model.Story
//	db.Table("a_story").Find(&Story)
//
//}

// 根据视频类型来获取全部视频
func GetVideoAllList(c *gin.Context) {
	videoAllList, count, err := services.GetVideoAllList(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video found", "data": videoAllList, "totalCount": count})
}

// 查询轮播图
func GetCarouselList(c *gin.Context) {
	carousel, err := services.GetCarouselList()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video found", "data": carousel})
}
