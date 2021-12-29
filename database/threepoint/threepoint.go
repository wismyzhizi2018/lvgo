package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// 变量后三个点表示将一个切片或数组变成一个一个的元素

	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// arr := []int{23, 454, 6}
	arrs := arr[0:]
	show1(arrs...)
}

func show1(x ...int) {
	for _, v := range x {
		fmt.Println(v)
	}
}

// 形参的参数前的三个点，表示可以传0到多个参数
func sum(args ...int) int {
	var out int
	for _, v := range args {
		out += v
	}
	return out
}
