package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked 1^1
	mutexWoken
	mutexStarving
	mutexStarvingF
	mutexStarvingS
	mutexWaiterShift = iota
)

func main() {
	fmt.Println("mutexLocked的值", mutexLocked)
	fmt.Println("mutexWoken的值", mutexWoken)
	fmt.Println("mutexStarving的值", mutexStarving)
	fmt.Println("mutexStarvingF的值", mutexStarvingF)
	fmt.Println("mutexStarvingS的值", mutexStarvingS)
	fmt.Println("mutexWaiterShift的值", mutexWaiterShift)
	go func() {
		fmt.Println("打印不?")

	}()
}
