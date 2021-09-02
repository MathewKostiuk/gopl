package main

import (
	"fmt"
	"unicode"
)

var b []byte
var run bool
var r []int

func main() {
	b = []byte("Hello    how many spaces are there   in this sentence   ?")
	run = false
	fmt.Printf("%v\n", b)

	for i := 0; i < len(b); i++ {
		if unicode.IsSpace(rune(b[i])) {
			run = unicode.IsSpace(rune(b[i+1]))
			if run {
				r = append(r, i)
				continue
			}

			if len(r) > 0 && !run {
				r = append(r, i)
				b = append(b[:r[0]], b[r[len(r)-1]:]...)
				r = r[:0]
			}
		}
	}
	fmt.Println(string(b[:]))
}
