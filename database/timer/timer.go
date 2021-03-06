package main

import (
	"time"

	"github.com/gookit/color"
)

// 类似js 中的setTimeout函数
func main() {
	// t := time.NewTicker(2 * time.Second)
	t1 := time.NewTimer(2 * time.Second)

	// time := <-t.C
	time2 := <-t1.C
	// color.Info.Println(time.Format("2006-01-01 15:04:05"))

	color.Danger.Println(time2.Format("2006-01-01 15:04:05"))
}
