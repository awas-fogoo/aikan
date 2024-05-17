package services

import (
	"github.com/gin-gonic/gin"
	"one/common"
	"one/model"
	"one/vo"
	_ "one/vo"
)

// 视频搜索
func VideoSearch(c *gin.Context) ([]vo.VideoMsg, int64, error) {
	type RequestBody struct {
		VideoName string `json:"videoName"`
		Page      int    `json:"page"`
		PageSize  int    `json:"pageSize"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	videoName := requestBody.VideoName
	db := common.DB
	var videoMsg []vo.VideoMsg
	// 执行模糊搜索并分页
	page := requestBody.Page
	pageSize := requestBody.PageSize
	//判断前端传参数是否正确
	if requestBody.VideoName == "" && requestBody.Page == 0 && pageSize < 1 {
		c.JSON(400, gin.H{"error": "请传入正确参数"})
	}
	offset := (page - 1) * pageSize
	var totalCount int64
	//查询搜索的总数，用来前端分页
	db.Table("videos").Where("title LIKE ?", "%"+videoName+"%").Count(&totalCount)
	//搜索出来的列表 指定分页
	db.Table("videos").Where("title LIKE ?", "%"+videoName+"%").Offset(offset).Limit(pageSize).Find(&videoMsg)
	return videoMsg, totalCount, nil
}

// 获取推荐视频
func GetRecommendVideoList(c *gin.Context) ([]model.Video, error) {
	var videoList []model.Video
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
	db.Where("story_id=?", storyId, "is_recommend", 1).Limit(20).Find(&videoList)
	//判断查询的是否为空
	if videoList == nil {
		c.JSON(200, gin.H{"message": "Video found", "data": "没有此类型的视频信息"})
	}
	return videoList, nil
}

// 根据视频类型来获取全部视频
func GetVideoAllList(c *gin.Context) ([]vo.VideoMsg, int64, error) {
	type RequestBody struct {
		StoryID  string `json:"storyID"`
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	db := common.DB
	var videoMsg []vo.VideoMsg
	// 执行模糊搜索并分页
	page := requestBody.Page
	pageSize := requestBody.PageSize
	//判断前端传参数是否正确
	if requestBody.StoryID == "" && requestBody.Page == 0 && pageSize < 1 {
		c.JSON(400, gin.H{"error": "请传入正确参数"})
	}
	offset := (page - 1) * pageSize
	var totalCount int64
	//查询搜索的总数，用来前端分页
	db.Table("videos").Where("story_id", requestBody.StoryID).Count(&totalCount)
	//搜索出来的列表 指定分页
	db.Table("videos").Where("story_id", requestBody.StoryID).Offset(offset).Limit(pageSize).Find(&videoMsg)
	return videoMsg, totalCount, nil
}
