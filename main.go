package main

import (
	"awesomeProject0511/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 封装路由 go run main.go
	r = routes.CollectRouter(r)

	// 启动HTTP服务,默认在0.0.0.0:8080启动服务
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
