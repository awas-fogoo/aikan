package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"one/common"
	"one/routes"
	"os"
)

func init() {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	InitConfig()
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 封装路由 go run main.go
	r = routes.CollectRouter(r)

	// 启动HTTP服务,默认在0.0.0.0:8080启动服务
	port := viper.GetString("server.port")
	fmt.Println(port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	common.InitDB()
	//common.InitCache()
	//common.InitMongoDB()
}
