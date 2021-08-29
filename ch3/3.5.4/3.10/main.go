package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	str := strings.Split(s, string('.'))
	n := len(str[0])

	if strings.Contains(str[0], string('.')) {
		n--
	}

	if n <= 3 {
		return str[0]
	}

	r := n % 3
	var m int

	for i, c := range str[0] {
		buf.WriteRune(c)
		if c == '-' {
			continue
		}

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

	if len(str) > 1 {
		for i, s := range str {
			if i != 0 {
				buf.WriteRune('.')
				buf.WriteString(s)
			}
		}
	}

	return buf.String()
}
