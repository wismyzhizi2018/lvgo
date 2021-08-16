package driver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gookit/color"
	"order/app/Common"
	"order/config"
)

var RedisDb *redis.Client

func InitRedis() {
	color.Debug.Println("尝试连接GoRedis...")

	rdbConfig := config.GetRedisConfig()

	RedisDb = redis.NewClient(&redis.Options{ // 连接服务
		Addr:     rdbConfig["Addr"],                        // string
		Password: rdbConfig["Password"],                    // string
		DB:       int(Common.StringToInt(rdbConfig["DB"])), // int
	})
	RedisPong, RedisErr := RedisDb.Ping(context.Background()).Result() // 心跳
	if RedisErr != nil {
		color.Danger.Println("Redis服务未运行。。。", RedisPong, RedisErr)
		//os.Exit(200)
	} else {
		color.Info.Println("GoRedis已连接 >>> ")
	}
}
