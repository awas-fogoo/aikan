package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/dto"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func LikeServer(vid string, uid uint) dto.RetStruct {
	rdb := common.InitCache()
	ctx := common.Ctx
	//likeVidIsExistRedis := LikeVidIsExistRedis(rdb, ctx, vid)
	//if likeVidIsExistRedis == 1 {
	//	return dto.RetStruct{
	//		Ret: true,
	//	}
	//}
	likeStatus := LikeVidStatus(rdb, ctx, vid, uid) // 点赞功能实现
	return dto.RetStruct{
		Ret: true,
		Data: gin.H{
			"like_status": likeStatus,
		},
	}
}

// LikeVidIsExistRedis 查看喜欢的vid是否存在redis里面
func LikeVidIsExistRedis(rdb *redis.Client, ctx context.Context, vid string) uint {
	vid = strconv.Itoa(1) + "_like_" + vid
	result, err := rdb.Exists(ctx, vid).Result()
	if err != nil {
		log.Panicln(err)
		return 0
	}
	return uint(result)
}

// LikeVidStatus 点赞&&取消点赞 key uid+_like_+vid
func LikeVidStatus(rdb *redis.Client, ctx context.Context, vid string, uid uint) bool {
	key := strconv.Itoa(int(uid)) + "_like_" + vid
	val, err := rdb.SIsMember(ctx, key, uid).Result()
	if err != nil {
		log.Panicln(err)
		return false
	}
	// 添加
	if val == false {
		_, errAdd := rdb.SAdd(ctx, key, uid).Result()
		if errAdd != nil {
			log.Panicln(errAdd)
		}
		return true
	} else {
		// 取消
		_, errCal := rdb.SRem(ctx, key, uid).Result()
		if errCal != nil {
			log.Panicln(errCal)
		}
		return false

	}
}
