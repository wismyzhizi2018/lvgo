package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

/**
*Context 使用原则
*1不要把Context放在结构体中，要以参数的方式传递
*2以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
*3给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
*4Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
*5Context是县城安全的，可以放心的在多个goroutine中传递
 */

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//WithValue传递元数据
	//    valueCtx:=context.WithValue(ctx,key,"【监控1】")
	go watch(context.WithValue(ctx, key, "【监控1】"))
	go watch(context.WithValue(ctx, key, "【监控2】"))
	go watch(context.WithValue(ctx, key, "【监控3】"))

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("name"), "监控退出，停止了...")
			return
		default:
			fmt.Println(ctx.Value("name"), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
