package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"awesomeProject0511/vo"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	db.Raw("SELECT cv.id,cv.cover,cv.hot,cv.title,cv.`desc`,cv.vid,cv.uid,cv.time,ui.`name`,ui.fans,ui.follow,ui.head_img,cn.click,cn.comment,cn.danmu,cn.like,cn.dislike,cn.collection,cn.share FROM channel_videos cv LEFT JOIN user_infos ui ON ui.uid=cv.uid LEFT JOIN channel_video_info_nums cn ON cn.uid=cv.uid").Scan(&channelListVos)
	// 获取视频相关推荐
	db.Model(&model.ChannelRecommend{}).Select("*").Limit(8).Scan(&upRecommendVos)
	res := dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"contentList": channelListVos,
			"upRecommend": upRecommendVos,
		},
	}

	// 创建视频信息缓存缓存
	createVideoNumCache := CreateVideoNumCache(rdb, ctx, vid, channelListVos, upRecommendVos)
	fmt.Printf("create redis cache %T\n", createVideoNumCache)
	return res
}

// CreateVideoNumCache 创建视频数据数量缓存
func CreateVideoNumCache(rdb *redis.Client, ctx context.Context, vid string, channelListVos []vo.ChannelListVo, upRecommendVos []vo.UpRecommendVo) bool {
	durationNum := 15 * time.Minute // 格式

	// 命名格式为vid+返回格式的json
	errStatus := rdb.Set(ctx, vid+"_status", "1", durationNum).Err()
	if errStatus != nil {
		panic(errStatus)
		return false
	}

	// 序列化两个map
	data, err := json.Marshal(channelListVos)
	data2, err2 := json.Marshal(upRecommendVos)

	if err != nil {
		fmt.Println("序列化_channelList失败")
	} else {
		errStatus := rdb.Set(ctx, vid+"_channelList", string(data), durationNum).Err()
		if errStatus != nil {
			panic(errStatus)
			return false
		}
	}

	if err2 != nil {
		fmt.Println("序列化_upRecommend失败")
	} else {
		errStatus := rdb.Set(ctx, vid+"_upRecommend", string(data2), durationNum).Err()
		if errStatus != nil {
			panic(errStatus)
			return false
		}
	}

	return true
}

// GetVideoNumCache 获取视频数据缓存
func GetVideoNumCache(rdb *redis.Client, ctx context.Context, vid string) dto.RetStruct {
	_, err := rdb.Get(ctx, vid+"_status").Result() // 判断vid的缓存是否已经存在
	if err == redis.Nil {
		fmt.Println("redis key does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		// 获取 channelList
		channelListVal, errStatus := rdb.Get(ctx, vid+"_channelList").Result()
		if errStatus != nil {
			panic(errStatus)
		}
		// 设置返回格式
		var channelListVos []vo.ChannelListVo
		err := json.Unmarshal([]byte(channelListVal), &channelListVos)
		if err != nil {
			fmt.Println("反序列化channelListVos失败")
		}

		// 获取 upRecommend
		upRecommendVal, errStatus2 := rdb.Get(ctx, vid+"_upRecommend").Result()
		if errStatus2 != nil {
			panic(errStatus2)
		}
		var upRecommendVos []vo.UpRecommendVo
		err2 := json.Unmarshal([]byte(upRecommendVal), &upRecommendVos)
		if err2 != nil {
			fmt.Println("反序列化upRecommendVos失败")
		}
		return dto.RetStruct{
			Ret: true,
			Data: gin.H{
				"contentList": channelListVos,
				"upRecommend": upRecommendVos,
			},
		}
	}
	return dto.RetStruct{Ret: false}
}
