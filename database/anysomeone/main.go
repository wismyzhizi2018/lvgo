package main

import "fmt"

func Print[T any](arr []T) error {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
	return nil
}

func main() {
	stars := []string{"Hello", "World", "Generics"}
	decks := []float64{3.14, 1.14, 1.618, 2.718}
	decks32 := []float32{3.14, 1.14, 1.618, 2.718}
	nums := []int{2, 4, 6, 8}

	_ = Print(stars)
	_ = Print(decks)
	_ = Print(nums)
	_ = Print(decks32)
}
