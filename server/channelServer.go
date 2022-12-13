package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetChannelService(vid string) dto.RetStruct {
	db := common.InitDB()
	defer db.Close()
	var channelListVos []vo.ChannelListVo
	var upRecommendVos []vo.UpRecommendVo
	var userInfoVo []vo.UserInfoVo
	var channelVideoInfoNumVo []vo.ChannelVideoInfoNumVo
	db.Model(&model.ChannelVideo{}).Where("vid = ?", vid).Scan(&channelListVos)
	db.Model(&model.ChannelRecommend{}).Select("*").Limit(8).Scan(&upRecommendVos)
	db.Model(&model.UserInfo{}).Where("uid = ?", channelListVos[0].Uid).Scan(&userInfoVo)
	db.Model(&model.ChannelVideoInfoNum{}).Where("vid = ?", vid).Scan(&channelVideoInfoNumVo)
	fmt.Println(channelVideoInfoNumVo)
	return dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"contentList": gin.H{
				"id":      channelListVos[0].ID,
				"imgUrl":  channelListVos[0].Cover,
				"hot":     channelListVos[0].Hot,
				"title":   channelListVos[0].Title,
				"desc":    channelListVos[0].Desc,
				"vid":     channelListVos[0].Vid,
				"uid":     channelListVos[0].Uid,
				"time":    channelListVos[0].CreatedAt.Format("2006-01-02 15:04:05"),
				"fans":    userInfoVo[0].Fans,
				"headUrl": userInfoVo[0].HeadImg,
				"name":    userInfoVo[0].Name,
				"see":     channelVideoInfoNumVo[0].Clicks,
				"danmu":   channelVideoInfoNumVo[0].Danmu,
				"collect": channelVideoInfoNumVo[0].Collects,
				"comment": channelVideoInfoNumVo[0].Comments,
				"like":    channelVideoInfoNumVo[0].Likes,
				"dislike": channelVideoInfoNumVo[0].Dislikes,
				"share":   channelVideoInfoNumVo[0].Shares,
			},
			"upRecommend": upRecommendVos,
		},
	}
}
