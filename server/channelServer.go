package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func GetChannelService(vid string) dto.RetStruct {
	var channelListVos []vo.ChannelListVo
	var upRecommendVos []vo.UpRecommendVo

	// 缓存临时数据
	rdb := common.InitCache()
	ctx := common.Ctx

	// 判断视频信息是否已经缓存
	getVideoNumCache := GetVideoNumCache(rdb, ctx, vid)
	if getVideoNumCache.Ret {
		return getVideoNumCache
	}

	db := common.InitDB()
	defer db.Close()
	// 获取视频信息
	db.Raw("SELECT cv.id,cv.cover,cv.hot,cv.title,cv.`desc`,cv.vid,cv.uid,cv.time,ui.`name`,ui.fans,ui.follow,ui.head_img,cn.click,cn.comment,cn.danmu,cn.like,cn.dislike,cn.collection,cn.share FROM channel_videos cv LEFT JOIN user_infos ui ON ui.uid=cv.uid LEFT JOIN channel_video_info_nums cn ON cn.uid=cv.uid WHERE cv.vid = " + vid).Scan(&channelListVos)
	if len(channelListVos) == 0 {
		return dto.RetStruct{
			Ret: true,
			Msg: "视频不存在",
		}
	}
	// 获取视频相关推荐
	db.Model(&model.ChannelRecommend{}).Select("*").Where("uid = ?", channelListVos[0].Uid).Limit(8).Scan(&upRecommendVos)
	res := dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"contentList": channelListVos,
			"upRecommend": upRecommendVos,
		},
	}

	// 创建视频信息缓存缓存
	createVideoNumCache := CreateVideoNumCache(rdb, ctx, vid, channelListVos, upRecommendVos)
	log.SetPrefix("[" + vid + "]")
	log.Println("create redis cache: ", createVideoNumCache)
	return res
}

// CreateVideoNumCache 创建视频数据数量缓存
func CreateVideoNumCache(rdb *redis.Client, ctx context.Context, vid string, channelListVos []vo.ChannelListVo, upRecommendVos []vo.UpRecommendVo) bool {
	durationNum := 15 * time.Minute // 格式

	// 命名格式为vid+返回格式的json
	// 序列化两个map
	data, err := json.Marshal(channelListVos)
	data2, err2 := json.Marshal(upRecommendVos)

	if err != nil {
		log.Println("序列化_channelList失败")
		return false
	} else {
		errStatus := rdb.Set(ctx, vid+"_channelList", string(data), durationNum).Err()
		if errStatus != nil {
			log.Panicln("serialize crete redis key false:", errStatus)
			return false
		}
	}

	if err2 != nil {
		log.Println("序列化_upRecommend失败")
		return false
	} else {
		errStatus := rdb.Set(ctx, vid+"_upRecommend", string(data2), durationNum).Err()
		if errStatus != nil {
			log.Panicln("serialize crete redis key false:", errStatus)
			return false
		}
	}

	return true
}

// GetVideoNumCache 获取视频数据缓存
func GetVideoNumCache(rdb *redis.Client, ctx context.Context, vid string) dto.RetStruct {
	val, err := rdb.Exists(ctx, vid+"_channelList").Result() // 判断vid_channelList的缓存是否已经存在
	if val == 0 {
		log.Println("redis key does not exist")
		return dto.RetStruct{
			Ret: false,
			Msg: "redis key does not exist",
		}
	} else if err != nil {
		log.Panicln(err)
	} else {
		// 获取 channelList
		channelListVal, errStatus := rdb.Get(ctx, vid+"_channelList").Result()
		if errStatus != nil {
			log.Panicln("serialize get redis key not exist:", errStatus)
		}
		// 设置返回格式
		var channelListVos []vo.ChannelListVo
		err := json.Unmarshal([]byte(channelListVal), &channelListVos)
		if err != nil {
			log.Println("反序列化channelListVos失败")
		}

		// 获取 upRecommend
		upRecommendVal, errStatus2 := rdb.Get(ctx, vid+"_upRecommend").Result()
		if errStatus2 != nil {
			log.Panicln(errStatus2)
		}
		var upRecommendVos []vo.UpRecommendVo
		err2 := json.Unmarshal([]byte(upRecommendVal), &upRecommendVos)
		if err2 != nil {
			log.Println("反序列化upRecommendVos失败")
		}
		return dto.RetStruct{
			Ret: true,
			Data: gin.H{
				"contentList": channelListVos,
				"upRecommend": upRecommendVos,
			},
		}
	}
	return dto.RetStruct{}
}
