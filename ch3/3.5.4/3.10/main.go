package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)
	if n <= 3 {
		return s
	}

	r := n % 3
	var m int

	for i, c := range s {
		buf.WriteRune(c)

		if r == 0 && i == 2 {
			buf.WriteByte(',')
			m = i
			continue
		}

		if i+1 == r || i == m+3 && i != n-1 {
			buf.WriteByte(',')
			m = i
		}
	}

	return buf.String()
}
