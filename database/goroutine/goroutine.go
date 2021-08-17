package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func trace(start, end byte) {
	for i := start; i < end; i++ {
		fmt.Printf("%c \n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(1) //限制只有一个逻辑处理器
	var wg sync.WaitGroup //用于等待所有协程都完成
	wg.Add(2)
	go func() {
		defer wg.Done() //程序退出的时候执行
		trace('A', 'Z')
	}()
	go func() {
		defer wg.Done() //程序退出的时候执行
		trace('a', 'z')
	}()
	wg.Wait() //等待所有协程的完成
	fmt.Println("结束")
}
