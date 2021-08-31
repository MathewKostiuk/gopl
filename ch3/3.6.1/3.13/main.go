package main

import "fmt"

const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)

func main() {
	fmt.Printf("KB=%v\n", KB)
	fmt.Printf("MB=%v\n", MB)
	fmt.Printf("GB=%v\n", GB)
	fmt.Printf("TB=%v\n", TB)
	fmt.Printf("PB=%v\n", PB)
	fmt.Printf("EB=%v\n", EB)
	fmt.Printf("ZB=%v\n", ZB)
	fmt.Printf("YB=%v\n", YB)
}
