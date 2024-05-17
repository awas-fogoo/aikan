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
	//
	db := common.DB
	//视频信息对象
	var video model.Video
	//对应视频信息链接对象
	var videoURL []model.VideoURL
	type RequestBody struct {
		VideoId int `json:"videoId"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//赋值
	videoId := requestBody.VideoId
	db.Table("videos").Where("id", videoId).First(&video)
	db.Table("video_urls").Where("video_id", videoId).Find(&videoURL)
	video.VideoURLs = videoURL
	c.JSON(200, gin.H{"message": "Video found", "data": video})
}

// 获取视频类型(1.电影 2.电视剧 3.综艺)
func GetVideoStory(c *gin.Context) {
	db := common.DB
	var Story []model.Story
	db.Table("a_story").Find(&Story)

}

// 视频搜索
func VideoSearch(c *gin.Context) {
	videosMsg, totalCount, err := services.VideoSearch(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video found", "data": videosMsg, "totalCount": totalCount})

}

// 根据视频类型来获取全部视频
func GetVideoAllList(c *gin.Context) {
	videoAllList, count, err := services.GetVideoAllList(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Video found", "data": videoAllList, "totalCount": count})
}
