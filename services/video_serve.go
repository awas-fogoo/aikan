package services

import (
	"github.com/gin-gonic/gin"
	"one/common"
	"one/model"
	"one/vo"
	_ "one/vo"
)

// 获取推荐视频
func GetRecommendVideoList(c *gin.Context) ([]vo.DetailMsg, error) {
	var videoList []vo.DetailMsg
	type RequestBody struct {
		StoryID int `json:"storyId"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

	}
	//赋值
	storyId := requestBody.StoryID
	//判断值是否为0
	if storyId == 0 {
		c.JSON(200, gin.H{"message": "Video found", "data": "传入信息有误"})
	}
	db := common.DB
	db.Table("details").Where("categories=?", storyId, "is_recommend", 1).Limit(20).Find(&videoList)
	//判断查询的是否为空
	if videoList == nil {
		c.JSON(200, gin.H{"message": "Video found", "data": "没有此类型的视频信息"})
	}
	return videoList, nil
}

// 根据视频类型来获取全部视频
func GetVideoAllList(c *gin.Context) ([]vo.DetailMsg, int64, error) {
	type RequestBody struct {
		CategoriesId string `json:"categoriesId"`
		Page         int    `json:"page"`
		PageSize     int    `json:"pageSize"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	db := common.DB
	var detailMsg []vo.DetailMsg
	// 执行模糊搜索并分页
	page := requestBody.Page
	pageSize := requestBody.PageSize
	//判断前端传参数是否正确
	if requestBody.CategoriesId == "" && requestBody.Page == 0 && pageSize < 1 {
		c.JSON(400, gin.H{"error": "请传入正确参数"})
	}
	offset := (page - 1) * pageSize
	var totalCount int64
	//查询搜索的总数，用来前端分页
	db.Table("details").Where("categories", requestBody.CategoriesId).Count(&totalCount)
	//搜索出来的列表 指定分页
	db.Table("details").Where("categories", requestBody.CategoriesId).Offset(offset).Limit(pageSize).Find(&detailMsg)
	return detailMsg, totalCount, nil
}

// 根据视频id获取
func GetVideoMsgByVideoId(c *gin.Context) (model.Detail, error) {
	//
	db := common.DB
	//视频信息对象
	var detail model.Detail
	//对应视频信息链接对象
	var video []model.Video
	type RequestBody struct {
		VideoId int `json:"videoId"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

	}
	//赋值
	videoId := requestBody.VideoId
	db.Table("details").Where("id", videoId).First(&detail)
	db.Table("video_urls").Where("video_id", videoId).Find(&video)
	detail.Videos = video
	//c.JSON(200, gin.H{"message": "Video found", "data": detail})
	return detail, nil
}

// 获取轮播图信息
func GetCarouselList() ([]model.Carousel, error) {
	db := common.DB
	var carousel []model.Carousel
	db.Table("carousels").Find(&carousel)
	return carousel, nil
}
