package main

import "fmt"

type coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

func main() {
	p := &coordinate{
		Lat: 0,
		Lng: 0,
	}
	fmt.Printf("Person =%#v", p)
	fmt.Println("\n")
	fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(32), "color show"))
}
