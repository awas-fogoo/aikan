package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"awesomeProject0511/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func LikeServer(vid string, uid uint) dto.RetStruct {
	db := common.InitDB()
	defer db.Close()
	var channel model.ChannelVideo
	db.Where("vid = ?", vid).First(&channel)
	//用户是否存在
	if channel.ID == 0 {
		return dto.RetStruct{
			Ret:  false,
			Code: 400,
			Msg:  "视频不存在",
		}
	}
	rdb := common.InitCache()
	ctx := common.Ctx
	//isLike := IsLike(rdb, ctx, vid, uid)
	likeStatus := LikeVidStatus(rdb, ctx, vid, uid) // 点赞功能实现
	cont := LikeIdCount(rdb, ctx, vid)              // 看某一视频的总点赞量
	//list := rdb.SMembers(ctx, videoUrl)
	//fmt.Println(list) // 所有点赞的人
	return dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"like": likeStatus,
			"cont": cont,
		},
	}
}

// IsLike 判断当前用户是否点赞
func IsLike(rdb *redis.Client, ctx context.Context, vid string, uid uint) bool {
	val, err := rdb.SIsMember(ctx, vid, uid).Result()
	if err != nil {
		panic(err)
	}
	if val == false {
		//fmt.Println("user don't like it")
		return false
	}
	return true

}

// LikeVidStatus 点赞&&取消点赞
func LikeVidStatus(rdb *redis.Client, ctx context.Context, vid string, uid uint) bool {
	val, err := rdb.SIsMember(ctx, vid, uid).Result()
	if err != nil {
		panic(err)
		return false
	}
	// 添加
	if val == false {
		_, errAdd := rdb.SAdd(ctx, vid, uid).Result()
		if errAdd != nil {
			panic(errAdd)
		}
		return true
	} else {
		// 取消
		_, errCal := rdb.SRem(ctx, vid, uid).Result()
		if errCal != nil {
			panic(errCal)
		}
		return false

	}
}

// LikeIdCount 获赞次数
func LikeIdCount(rdb *redis.Client, ctx context.Context, vid string) int64 {
	val, err := rdb.SCard(ctx, vid).Result()
	if err != nil {
		panic(err)
	}
	return val
}
