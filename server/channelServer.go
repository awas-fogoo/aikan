package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func GetChannelService(vid string) dto.RetStruct {
	db := common.InitDB()
	defer db.Close()
	var channelListVos []vo.ChannelListVo
	db.Model(&model.ChannelVideo{}).Where("vid = ?", vid).Scan(&channelListVos)
	return dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"contentList": gin.H{
				"id":     channelListVos[0].ID,
				"imgUrl": channelListVos[0].Cover,
				"hot":    channelListVos[0].Hot,
				"title":  channelListVos[0].Title,
				"desc":   channelListVos[0].Desc,
				"vid":    channelListVos[0].Vid,
				"uid":    channelListVos[0].Uid,
				"time":   channelListVos[0].CreatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	}
}
