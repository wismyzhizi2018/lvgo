package main

import (
	"fmt"
	"sync"
)

type OutOrderInfo struct {
	title string
}

func main() {
	// 执行的 这里要注意  需要指针类型传入  否则会异常
	wg := &sync.WaitGroup{}
	// 并发控制 10
	limiter := make(chan struct{}, 20)
	defer close(limiter)

	response := make(chan *OutOrderInfo, 20)
	wgResponse := &sync.WaitGroup{}
	// var result []string
	// 处理结果 接收结果
	go func() {
		wgResponse.Add(1)
		for rc := range response {
			fmt.Println(rc.title)
		}
		wgResponse.Done()
	}()
	// 发送请求
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//	cmd := &LazadaInfo{AccessToken: token, OrderId: orderId, Country: country, OutInfo: outRow, Wg: &wg, Ch: ch}
		// 并发控制 20
		limiter <- struct{}{}
		go pushLazadaGetOrderItems(fmt.Sprintf("Hellow GoRoutine! %d", i), wg, response, limiter)
	}
	//发送任务
	wg.Wait()
	fmt.Println("发送完毕")
	close(response) // 关闭 并不影响接收遍历
	// 处理接收结果
	wgResponse.Wait()
	fmt.Println("请求结束")
}

func pushLazadaGetOrderItems(AccessToken string, Wg *sync.WaitGroup, response chan *OutOrderInfo, limiter chan struct{}) {
	// 计数器-1
	defer Wg.Done()
	// AccessToken := _this.AccessToken
	// OrderId := _this.OrderId
	// Country := _this.Country
	// OutInfo := _this.OutInfo
	out := &OutOrderInfo{title: AccessToken}
	if AccessToken != "" {
		response <- out
	} else {
		response <- &OutOrderInfo{}
	}
	// 释放一个并发
	<-limiter
}
