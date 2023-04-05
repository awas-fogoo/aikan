package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/util"
	"github.com/gin-gonic/gin"
)

func AddCommentVideoServer(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	user, _ := c.Get("user")
	userDto := dto.ToUserDTO(user.(model.User))
	content := c.PostForm("content")
	userID := userDto.ID
	videoID := util.StringToUint(c.PostForm("video_id"))
	parentCommentID := util.StringToUint(c.PostForm("parent_id"))
	if len(content) < 0 && videoID < 0 && parentCommentID < 0 {
		c.JSON(0, dto.Error(-1, "parameter error"))
		return
	}
	var parentID *uint
	var ancestorID uint
	var level uint

	if parentCommentID == 0 {
		parentID = nil
		ancestorID = 0
		level = 0
	} else {
		var parentComment model.Comment
		var commentRelation model.CommentRelation
		if err := db.Where("id = ?", parentCommentID).First(&parentComment).Error; err != nil {
			// 父评论不存在，这里可以根据需求决定如何处理
			c.JSON(0, dto.Error(-1, "parent id does not exist"))
		} else {
			parentID = &parentCommentID
			if parentComment.ParentID == nil {
				ancestorID = parentComment.ID
			} else {
				ancestorID = *parentComment.ParentID
			}
			level = commentRelation.Level + 1
		}
	}

	// 创建一条新评论
	comment := model.Comment{
		Content:  content,
		UserID:   userID,
		VideoID:  videoID,
		ParentID: parentID,
	}
	db.Create(&comment)

	// 创建评论关系
	if parentCommentID != 0 {
		commentRelation := model.CommentRelation{
			AncestorID:   ancestorID,
			DescendantID: comment.ID,
			Level:        level,
		}
		db.Create(&commentRelation)
	}
	c.JSON(0, dto.Success("comment success"))
	return
}
