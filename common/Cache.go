package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Ctx = context.Background()

func InitCache() *redis.Client {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: viper.GetString("redis.password"),
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		panic("faild to connect redis ,err :" + err.Error())
	} else {
		return client
	}
}
