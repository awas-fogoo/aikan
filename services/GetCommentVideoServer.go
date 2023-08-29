package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GerCommentVideoServer(c *gin.Context) {
	db := common.DB
	videoID := c.Query("vid")
	// 一级评论
	var comments []model.Comment
	db.Preload("User").Where("video_id = ? AND parent_id is null", videoID).Find(&comments)
	fmt.Printf("有%d条一级评论", len(comments))
	for _, c := range comments {
		fmt.Println("一级评论：", c.Content)
		// 二级评论
		var children []model.Comment
		db.Preload("User").Where("parent_id = ?", c.ID).Find(&children)
		for _, cc := range children {
			fmt.Println("    二级评论：", cc.Content)
			// 三级评论
			var grandchildren []model.Comment
			db.Preload("User").Where("parent_id = ?", cc.ID).Find(&grandchildren)
			for _, gcc := range grandchildren {
				fmt.Println("        三级评论：", gcc.Content)
			}
		}
	}
	c.JSON(200, dto.Success("get comment content success"))
	return
}
