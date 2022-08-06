package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/model"
	"awesomeProject0511/response"
	"awesomeProject0511/vo"
	"github.com/gin-gonic/gin"
)

func GetSwiperListService() response.ResStruct {
	db := common.InitDB()
	defer db.Close()

	var swiperListVos []vo.SwiperListVo
	db.Model(&model.SwiperList{}).Select("id, uid, img_url, video_id").Scan(&swiperListVos)
	return response.ResStruct{
		Ret: true,
		Data: gin.H{
			"popRecommend":  "首页推荐",
			"classicReview": "大家常看",
			"swiperList":    swiperListVos,
		},
	}
}
