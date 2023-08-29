package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

func SearchVideoServer(c *gin.Context) {
	db := common.DB
	q := c.Query("q")

	if len(q) >= 100 || len(q) < 0 {
		c.JSON(200, dto.Error(4000, "搜索长度不是正确的"))
		return
	}

	var videos []vo.SearchVideoVo

	// 分页，默认1-10
	page := util.StringToUint(c.DefaultQuery("page", "1"))
	perPage := util.StringToUint(c.DefaultQuery("perPage", "10"))

	// 排序 默认根据创建时间的 降序
	sortField := c.DefaultQuery("sortField", "created_at")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	// 筛选
	durationMin := c.DefaultQuery("durationMin", "0")
	durationMax := c.DefaultQuery("durationMax", "999999")

	location := c.DefaultQuery("location", "CN")

	db.Table("videos").
		Select("videos.*, GROUP_CONCAT(tags.name SEPARATOR ', ') as tags").
		Joins("JOIN video_tags ON videos.id = video_tags.video_id").
		Joins("JOIN tags ON tags.id = video_tags.tag_id").
		Where("videos.deleted_at IS NULL AND duration >= ? AND duration <= ? AND (videos.title LIKE ? OR videos.description LIKE ? OR tags.name LIKE ?)", durationMin, durationMax, "%"+q+"%", "%"+q+"%", "%"+q+"%").
		Group("videos.id").
		Order(sortField + " " + sortOrder).
		Offset((page - 1) * perPage).
		Limit(perPage).
		Scan(&videos)

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(200, dto.Success(videos))
		return
	}

	// 创建缓存
	records := make(chan model.SearchRecord)
	// 开启并发
	go storeRecords(db, records)

	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	record := model.SearchRecord{
		UserID:   claims.UserId,
		Keyword:  q,
		Location: location,
	}
	records <- record
	time.Sleep(time.Second)
	var hotKeywords []struct {
		Keyword string
		Count   int
	}
	db.Table("search_records").Select("keyword, count(keyword) as count").Group("keyword").Order("count desc").Limit(10).Scan(&hotKeywords)
	c.JSON(200, dto.Success(videos))
}

func storeRecords(db *gorm.DB, records chan model.SearchRecord) {
	for record := range records {
		db.Create(&record)
	}
}
