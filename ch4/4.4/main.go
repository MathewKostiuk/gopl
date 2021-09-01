package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	p := rotate(a[:], 2)
	s := rotate(a[:], -2)
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(s)
}

// Rotate in a single pass
func rotate(s []int, n int) []int {
	r := make([]int, len(s), cap(s))

	for i := range s {
		if i+n >= len(s) {
			r[i] = s[i+n-len(s)]
			continue
		}

		if i+n < 0 {
			r[i] = s[i+n+len(s)]
			continue
		}
		r[i] = s[i+n]
	}

	return r
}
