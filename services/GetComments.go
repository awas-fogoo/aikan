package services

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
)

type CommentResult struct {
	Content   string `json:"content"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	CreatedAt int64  `json:"created_at"`
}

func (s VideoService) GetComments(vid string) ([]CommentResult, error) {
	// 连接数据库
	db := common.DB

	// 查询评论列表
	// 一级评论
	var comments []model.Comment
	db.Preload("User").Where("video_id = ? AND parent_id is null", vid).Limit(10).Find(&comments)
	var results []CommentResult
	for _, c := range comments {
		result := CommentResult{
			Content:   c.Content,
			AvatarUrl: c.User.AvatarUrl,
			Nickname:  c.User.Nickname,
			CreatedAt: c.CreatedAt.Unix(),
		}
		results = append(results, result)
	}
	return results, nil

}
