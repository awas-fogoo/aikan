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
	db.Model(&model.ChannelList{}).Where("vid = ?", vid).Scan(&channelListVos)
	return dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"time":        channelListVos[0].CreatedAt.Format("2006-01-02 15:04:05"),
			"contentList": channelListVos,
		},
	}
}
