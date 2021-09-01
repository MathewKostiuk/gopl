package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	p := rotate(a[:], 4)
	s := rotate(a[:], 2)
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(s)

	r := rotateSlice(a[:], 4)
	fmt.Println(r)
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

func rotateSlice(s []int, n int) []int {
	s1 := s[n:]
	s2 := s[:n]
	r := append(s1, s2...)
	return r
}
