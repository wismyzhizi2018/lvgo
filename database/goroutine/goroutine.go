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
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		trace('A', 'Z')
	}()
	go func() {
		defer wg.Done()
		trace('a', 'z')
	}()
	wg.Wait()
	fmt.Println("结束")
}
