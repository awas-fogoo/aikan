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
	"strconv"
	"time"
)

func LikeServer(uid uint, vid string) dto.RetStruct {
	db := common.InitDB()
	rdb := common.InitCache()
	ctx := common.Ctx
	rid := strconv.Itoa(int(uid)) + "_relation_" + vid
	videoRelationsRedis := GetVideoRelationsCache(rdb, ctx, rid)
	if !videoRelationsRedis {
		res := UpdateVideoRelationsCache(rdb, ctx, rid)
		return res
	}
	// create
	// write redis data
	var videoRelationVos []vo.VideoRelationVo
	db.Model(&model.VideoRelation{}).Where("uid = ? and vid = ?", uid, vid).Scan(&videoRelationVos)
	if len(videoRelationVos) == 0 {
		u := model.VideoRelation{
			Uid: uid, Vid: vid,
		}
		if err := db.Create(&u).Error; err != nil {
			log.Println("插入失败", err)
		}
		db.Model(&model.VideoRelation{}).Where("uid = ? and vid = ?", uid, vid).Scan(&videoRelationVos)
	}
	cache := CreateVideoRelationsCache(rdb, ctx, rid, videoRelationVos)
	if !cache {
		log.Panicln("CreateVideoRelationsCache ---false")
		return dto.RetStruct{}
	}

	// 定时提交redis
	go func() {
		for true {
			// 创建一个 Timer
			myT := time.NewTimer(1 * time.Second)
			//从通道中读取数据，若读取得到，说明时间到了
			val, err := rdb.TTL(ctx, rid).Result()
			if val == 1*time.Second {
				getVal, err := rdb.Get(ctx, rid).Result()
				if err != nil {
					log.Panicln("time get submit :", err)
				}
				err2 := json.Unmarshal([]byte(getVal), &videoRelationVos)
				if err2 != nil {
					log.Println("反序列化videoRelationVos失败")
				}
				var count int
				db.Model(&model.VideoRelation{}).Where("vid = ? and uid = ?", vid, uid).Update("is_like", videoRelationVos[0].IsLike)
				db.Model(&model.VideoRelation{}).Where("vid = ? and is_like = ?", vid, true).Count(&count)
				db.Model(&model.ChannelVideoInfoNum{}).Where("vid = ?", vid).Update("like", count)
			} else if err != nil {
				log.Panicln("time exist submit :", err)
			}
			<-myT.C
			if val == -2 {
				return
			}
		}
	}()
	return UpdateVideoRelationsCache(rdb, ctx, rid)
}

func UpdateVideoRelationsCache(rdb *redis.Client, ctx context.Context, rid string) dto.RetStruct {
	// 获取 rid
	videoRelationVal, errStatus := rdb.Get(ctx, rid).Result()
	if errStatus != nil {
		log.Panicln("serialize get redis key not exist:", errStatus)
		return dto.RetStruct{}
	}
	// 设置返回格式
	var videoRelationVos []vo.VideoRelationVo
	err := json.Unmarshal([]byte(videoRelationVal), &videoRelationVos)
	if err != nil {
		log.Println("反序列化videoRelationVos失败")
		return dto.RetStruct{}
	}
	if videoRelationVos[0].IsLike {
		videoRelationVos[0].IsLike = false
		CreateVideoRelationsCache(rdb, ctx, rid, videoRelationVos)
	} else if !videoRelationVos[0].IsLike {
		videoRelationVos[0].IsLike = true
		CreateVideoRelationsCache(rdb, ctx, rid, videoRelationVos)
	}
	return dto.RetStruct{
		Ret:  true,
		Data: gin.H{"res": videoRelationVos},
		Code: 200,
		Msg:  "like succ",
	}
}
func GetVideoRelationsCache(rdb *redis.Client, ctx context.Context, rid string) bool {
	val, err := rdb.Exists(ctx, rid).Result()
	if val == 0 {
		return true
	} else if err != nil {
		log.Panicln(err)
		return false
	}
	return false
}
func CreateVideoRelationsCache(rdb *redis.Client, ctx context.Context, rid string, videoRelationVo []vo.VideoRelationVo) bool {
	data, err := json.Marshal(videoRelationVo)
	if err != nil {
		log.Println("序列化_channelList失败")
		return false
	} else {
		val, err := rdb.Exists(ctx, rid).Result() // 判断rid的缓存是否已经存在
		if val == 0 {
			errStatus := rdb.Set(ctx, rid, string(data), 3*time.Second).Err()
			if errStatus != nil {
				log.Panicln("serialize crete redis key false:", errStatus)
				return false
			}
		} else if err != nil {
			log.Panicln(err)
		} else {
			rdb.Del(ctx, rid)
			errStatus := rdb.Set(ctx, rid, string(data), 3*time.Second).Err()
			if errStatus != nil {
				log.Panicln("serialize crete redis key false:", errStatus)
				return false
			}
		}

	}
	return true
}
