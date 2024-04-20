package common

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MDBMux sync.Mutex
	MDB    *mongo.Client // MongoDB 客户端
)

func InitMongoDB() {
	host := viper.GetString("mongo.host")
	port := viper.GetString("mongo.port")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")

	// 构建 MongoDB 连接 URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	clientOptions := options.Client().ApplyURI(uri)

	// 创建一个新的 MongoDB 客户端
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Printf("Failed to create MongoDB client %v", err)
		return
	}

	// 检查与数据库的连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("连接到 MongoDB 失败：%v", err)
	} else {
		log.Printf("Successfully connected to MongoDB")
		// 如果连接成功，则将新的客户端赋值给 MDB 变量
		MDBMux.Lock()
		MDB = client
		MDBMux.Unlock()
	}
}
