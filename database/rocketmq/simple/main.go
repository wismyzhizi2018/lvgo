package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// 延迟消息
func main() {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"61.171.40.222:9876"}))
	if err != nil {
		panic("生成producer失败")
	}

	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}
	for i := 0; i < 100; i++ {
		message := primitive.NewMessage("hellomq", []byte("this is imooc =>"+strconv.Itoa(i)))
		num := rand.Intn(60)
		message.WithDelayTimeLevel(num)
		res, err := p.SendSync(context.Background(), message)
		if err != nil {
			fmt.Printf("发送失败: %s\n", err)
		} else {
			fmt.Printf("发送成功: %s\n", res.String())
		}
	}
	if err = p.Shutdown(); err != nil {
		panic("关闭producer失败")
	}
}
