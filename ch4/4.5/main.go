package main

import "fmt"

var slice []string

func main() {
	slice = []string{"hello", "hello", "hell", "pillow", "hell", "foo", "foo"}
	adjacentDuplicateEliminator()
	fmt.Printf("s=%v\t&s=%p\n", slice, slice)
}

// Eliminates adjacent duplicates in a []string slice
func adjacentDuplicateEliminator() {
	for i := 0; i < len(slice); i++ {
		if slice[i] == slice[i+1] {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
}
