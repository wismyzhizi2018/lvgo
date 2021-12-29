package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"61.171.40.222:9876"}),
		consumer.WithGroupName("mxshop"),
	)
	_ = c.Subscribe("hellomq", consumer.MessageSelector{}, callback)
	_ = c.Start()
	time.Sleep(time.Second * 3600)
}

func callback(cxt context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i, message := range messages {
		fmt.Printf("Message=%d ===%s \r\n", i, message)
	}
	return consumer.ConsumeSuccess, nil
}
