package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
)

func LikeServer(vid string, uid uint) dto.RetStruct {
	db := common.InitDB()
	defer db.Close()
	var channel model.ChannelList
	db.Where("vid = ?", vid).First(&channel)
	//用户是否存在
	if channel.ID == 0 {
		return dto.RetStruct{
			Ret:  false,
			Code: 400,
			Msg:  "视频不存在",
		}
	}

	return dto.RetStruct{
		Ret:  true,
		Data: nil,
		Code: 200,
		Msg:  "点赞成功",
	}
}
