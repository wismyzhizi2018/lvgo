package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"time"
)

type TestListener struct {
}

func (t TestListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行1111")
	time.Sleep(time.Second * 2)
	fmt.Println("开始执行结束")
	return primitive.UnknowState
}

func (t TestListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("开始执行回查")
	time.Sleep(time.Second * 3)
	fmt.Println(ext.MsgId)
	return primitive.UnknowState
}

func NewTestListener() *TestListener {
	return &TestListener{}
}

func main() {
	p, _ := rocketmq.NewTransactionProducer(
		NewTestListener(),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"61.171.40.222:9876"})),
		producer.WithRetry(1),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
		os.Exit(1)
	}
	topic := "test"
	for i := 0; i < 2; i++ {
		res, err := p.SendMessageInTransaction(context.Background(),
			primitive.NewMessage(topic, []byte("Hello RocketMQ again "+strconv.Itoa(i))))

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	time.Sleep(5 * time.Minute)
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
