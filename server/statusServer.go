package server

import (
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func StatusServer(db *gorm.DB, vid string, uid uint) dto.RetStruct {
	var count int64
	db.Model(&model.ChannelLiked{}).Where("vid = ? and status = ?", vid, 1).Count(&count)
	// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)
	result := dto.RetStruct{
		Ret:  true,
		Code: 200,
		Data: gin.H{
			"like":  true,
			"count": count,
		},
		Msg: "OK",
	}

	_, status := IsMysqlLike(db, vid, uid)
	//验证是否已经点赞
	if status == 0 {
		result.Ret = true
		result.Code = 3011
		result.Data = gin.H{
			"like":  false,
			"count": count,
		}
		result.Msg = "已经点赞过了哦"
		return result
	}
	if status == -1 {
		liked := model.ChannelLiked{
			Uid:    uid,
			Vid:    vid,
			Status: true,
		}
		db.Create(&liked)
	} else {
		db.Model(&model.ChannelLiked{}).Where("uid = ? AND vid = ?", uid, vid).Update("status", true)
	}
	return result
}

func IsMysqlLike(db *gorm.DB, vid string, uid uint) (bool, int) {
	var liked model.ChannelLiked
	db.Where("uid = ? and vid = ?", uid, vid).First(&liked)
	//不存在返回-1，存在但是没有点赞返回1，已经点赞返回0
	if liked.ID == 0 {
		return false, -1
	} else if !liked.Status {
		return false, 1
	}
	return true, 0
}
