package main

import "github.com/gookit/color"

func addr() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAddr func(int) (int, iAddr)

func add2(base int) iAddr {
	return func(i int) (int, iAddr) {
		return base + i, add2(base + i)
	}
}
func main() {
	a := addr()
	for i := 0; i < 100; i++ {
		color.Redln(a(i))
	}
	b := add2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		color.Debug.Printf("0 + 1 +... + %d=%d\n", i, s)
	}
}
