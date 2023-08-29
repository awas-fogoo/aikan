package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math"
	"time"
)

// GetPopularVideos 获取观看次数最多的前N个视频
func getPopularVideos(db *gorm.DB, n int) ([]vo.VideoHomeVo, error) {
	var videos []model.Video
	var videoVos []vo.VideoHomeVo
	err := db.Order("views desc").Limit(n).Find(&videos).Scan(&videoVos).Error
	if err != nil {
		return nil, err
	}
	return videoVos, nil
}
func calculateVideoWeight(likes, comments, collections, danmaku, views uint, createdAt time.Time) float64 {
	// 根据具体应用需要自行调整参数的比例值
	const (
		likesWeight       = 1.0
		commentsWeight    = 0.5
		collectionsWeight = 0.3
		danmakuWeight     = 0.2
		viewsWeight       = 0.1
		timeWeight        = 0.5
	)

	// 计算一个视频的热度得分
	score := float64(likes)*likesWeight + float64(comments)*commentsWeight + float64(collections)*collectionsWeight + float64(danmaku)*danmakuWeight + float64(views)*viewsWeight

	// 计算视频的时间衰减因子
	now := time.Now()
	var ageFactor float64
	ageHours := now.Sub(createdAt).Hours() + 2 // 将ageHours加上一个小时来避免ageHours等于0的情况
	if ageHours <= 0 {
		ageFactor = 1.0
	} else {
		ageFactor = 1.0 / math.Pow(ageHours+2, timeWeight)
	}

	// 计算视频权重值
	weight := score * ageFactor

	// 避免NaN和Inf
	if math.IsNaN(weight) || math.IsInf(weight, 0) {
		weight = 0.0
	}

	return weight
}
func GetHotVideoServer(c *gin.Context) {
	db := common.DB
	videos, err := getPopularVideos(db, 4)
	if err != nil {
		log.Println(err)
	}
	// videos即为获取的热门视频列表
	c.JSON(200, dto.Success(videos))
}
