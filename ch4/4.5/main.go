package main

import "fmt"

func main() {
	s := [...]string{"hello", "hello", "hell", "pillow", "hell", "foo", "foo"}
	fmt.Println(s)
	adjacentDuplicateEliminator(s[:])
	fmt.Println(s)
}

// Eliminates adjacent duplicates in a []string slice
func adjacentDuplicateEliminator(s []string) {
	for i, r := range s {
		if i == len(s)-1 {
			break
		}
		if r == s[i+1] {
			s[i] = ""
		}
	}
}
