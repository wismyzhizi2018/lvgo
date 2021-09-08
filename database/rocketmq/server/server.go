package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type TestListener struct {
	localTrans       *sync.Map
	transactionIndex int32
}

func NewTestListener() *TestListener {
	return &TestListener{localTrans: new(sync.Map)}
}

func (t TestListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {
	nextIndex := atomic.AddInt32(&t.transactionIndex, 1)
	fmt.Printf("nextIndex: %v for transactionID: %v\n", nextIndex, message.TransactionId)
	status := nextIndex % 3
	t.localTrans.Store(message.TransactionId, primitive.LocalTransactionState(status+1))
	fmt.Printf("dl")
	return primitive.UnknowState
}

func (t TestListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Printf("%v msg transactionID : %v\n", time.Now(), ext.TransactionId)
	v, existed := t.localTrans.Load(ext.TransactionId)
	if !existed {
		fmt.Printf("unknow msg: %v, return Commit", ext)
		return primitive.CommitMessageState
	}
	state := v.(primitive.LocalTransactionState)
	switch state {
	case 1:
		fmt.Printf("checkLocalTransaction COMMIT_MESSAGE: %v\n", ext)
		return primitive.CommitMessageState
	case 2:
		fmt.Printf("checkLocalTransaction ROLLBACK_MESSAGE: %v\n", ext)
		return primitive.RollbackMessageState
	case 3:
		fmt.Printf("checkLocalTransaction unknow: %v\n", ext)
		return primitive.UnknowState
	default:
		fmt.Printf("checkLocalTransaction default COMMIT_MESSAGE: %v\n", ext)
		return primitive.CommitMessageState
	}
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
	for i := 0; i < 10; i++ {
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
