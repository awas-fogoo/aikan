package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"sync" // 导入 sync 包，用于确保线程安全
)

var (
	RDB    *redis.Client
	RDBMux sync.Mutex // 用于保证 RDB 的线程安全
)

func InitCache() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,   // 使用默认数据库
		PoolSize: 100, // 连接池大小
	})

	// 为这个操作使用特定的上下文
	ctx := context.TODO()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("Failed connected to Redis: %v", err)
	} else {
		log.Printf("Successfully connected to Redis")
		// 如果连接成功，则将新的客户端赋值给 RDB 变量
		RDBMux.Lock()
		RDB = client
		RDBMux.Unlock()
	}
}
