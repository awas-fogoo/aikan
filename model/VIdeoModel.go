package model

import "github.com/jinzhu/gorm"

/*
	like 在UserLike 更新，创建，删除之前，对video表中的likes字段进行递增与递减
*/

// BeforeCreate 创建点赞记录时，同时更新视频表中的点赞数量字段
func (like *UserLike) BeforeCreate(tx *gorm.DB) (err error) {
	// 更新视频表中的点赞数量字段
	err = tx.Model(&Video{}).Where("id = ?", like.VideoID).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
	return
}

// BeforeUpdate 创建点赞记录时，同时更新视频表中的点赞数量字段
func (like *UserLike) BeforeUpdate(tx *gorm.DB) (err error) {
	// 更新视频表中的点赞数量字段
	err = tx.Model(&Video{}).Where("id = ?", like.VideoID).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
	return
}

// AfterDelete 删除点赞记录时，同时更新视频表中的点赞数量字段
func (like *UserLike) AfterDelete(tx *gorm.DB) (err error) {
	// 更新视频表中的点赞数量字段
	err = tx.Model(&Video{}).Where("id = ?", like.VideoID).UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error
	return
}

/*
	Collection 在UserCollection 更新，创建，删除之前，对video表中的collections字段进行递增与递减
*/

// BeforeCreate 创建点赞记录时，同时更新视频表中的点赞数量字段
func (collection *UserCollection) BeforeCreate(tx *gorm.DB) (err error) {
	// 更新视频表中的点赞数量字段
	err = tx.Model(&Video{}).Where("id = ?", collection.VideoID).UpdateColumn("collections", gorm.Expr("collections + ?", 1)).Error
	return
}

// BeforeUpdate 创建点赞记录时，同时更新视频表中的点赞数量字段
func (collection *UserCollection) BeforeUpdate(tx *gorm.DB) (err error) {
	// 更新视频表中的点赞数量字段
	err = tx.Model(&Video{}).Where("id = ?", collection.VideoID).UpdateColumn("collections", gorm.Expr("collections + ?", 1)).Error
	return
}

// AfterDelete 删除点赞记录时，同时更新视频表中的点赞数量字段
func (collection *UserCollection) AfterDelete(tx *gorm.DB) (err error) {
	// 更新视频表中的点赞数量字段
	err = tx.Model(&Video{}).Where("id = ?", collection.VideoID).UpdateColumn("collections", gorm.Expr("collections - ?", 1)).Error
	return
}
