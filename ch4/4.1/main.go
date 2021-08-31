package main

import (
	"crypto/sha256"
	"fmt"
)

var count = 0

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	countDifferences(c1, c2)

	fmt.Printf("%d\n", count)
}

func countDifferences(c1, c2 [32]byte) {
	for i, b := range c1 {
		fmt.Printf("%v\t%v\n", b, c2[i])
		if b != c2[i] {
			count++
		}
	}
}
