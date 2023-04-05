package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/util"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func SearchVideoServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	q := c.Query("q")
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
	c.JSON(0, dto.Success(videos))
}
