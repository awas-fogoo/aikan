package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加剧集的处理函数  /episodes

func AddEpisode(c *gin.Context) {
	// 解析请求中的参数
	videoIDStr := c.PostForm("video_id")
	seasonNumStr := c.PostForm("season_num")
	episodeNumStr := c.PostForm("episode_num")
	title := c.PostForm("title")
	description := c.PostForm("description")
	url := c.PostForm("url")
	coverUrl := c.PostForm("cover_url")
	durationStr := c.PostForm("duration")

	// 解析 videoID 为 uint 类型
	videoID, err := strconv.ParseUint(videoIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	// 解析 seasonNum 为 uint 类型
	seasonNum, err := strconv.ParseUint(seasonNumStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season number"})
		return
	}

	// 解析 episodeNum 为 uint 类型
	episodeNum, err := strconv.ParseUint(episodeNumStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid episode number"})
		return
	}

	// 解析 duration 为 float64 类型
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration"})
		return
	}

	// 查询对应的电视剧
	// 连接数据库
	db := common.DB
	var video model.Video
	if err := db.First(&video, videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	// 检查季数是否已存在，如果不存在则创建新的季数
	var season model.Season
	if err := db.Where("video_id = ? AND season_num = ?", videoID, seasonNum).First(&season).Error; err != nil {
		season = model.Season{
			VideoID:   uint(videoID),
			SeasonNum: uint(seasonNum),
			Episodes:  make([]model.Episode, 0),
		}
	}

	// 创建剧集对象
	episode := model.Episode{
		SeasonID:    season.ID,
		EpisodeNum:  uint(episodeNum),
		Title:       title,
		Description: description,
		Url:         url,
		CoverUrl:    coverUrl,
		Duration:    duration,
	}

	// 将剧集添加到季数的剧集列表中
	season.Episodes = append(season.Episodes, episode)

	// 保存剧集和季数到数据库
	if err := db.Save(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save episode"})
		return
	}

	if err := db.Save(&season).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save season"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Episode added successfully"})
}
