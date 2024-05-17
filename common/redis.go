package common

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 访问密码
		DB:       0,                // 使用的数据库编号
	})

	// 测试连接
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

func SaveSearchHistory(userID string, keyword string) error {
	// 设置 Redis key，用于存储搜索历史
	key := "search_history:" + userID

	// 获取当前时间
	now := time.Now()

	// 存储搜索历史，并设置过期时间
	err := redisClient.ZAdd(context.Background(), key, &redis.Z{
		Score:  float64(now.Unix()),
		Member: keyword,
	}).Err()
	if err != nil {
		return err
	}

	// 设置搜索历史的过期时间（假设为一周）
	_, err = redisClient.Expire(context.Background(), key, 7*24*time.Hour).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetSearchHistory(userID string) ([]string, error) {
	// 设置 Redis key，用于存储搜索历史
	key := "search_history:" + userID

	// 获取搜索历史
	history, err := redisClient.ZRevRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return history, nil
}
