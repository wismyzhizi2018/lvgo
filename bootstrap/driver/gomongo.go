package driver

import (
	"context"

	"github.com/gookit/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongodbClient *mongo.Client

func InitMongo() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://192.168.0.252:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		color.Danger.Println("连接到MongoDB失败，MongoDB功能将不可用。。。", err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		color.Danger.Println("检查连接失败，MongoDB功能将不可用。。。", err)
	}
	MongodbClient = client
	color.Info.Println("Mongodb已连接 >>>")
}
