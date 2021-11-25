package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}
	b := false
	//优化
	//因为空结构体变量的内存占用大小为0，而bool类型内存占用大小为1，这样可以更加最大化利用我们服务器的内存空间
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
