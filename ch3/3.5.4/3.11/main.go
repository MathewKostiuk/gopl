package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(anagram(args[0], args[1]))
}

// Anagram checks if two strings are anagrams
// (they contain the same letters in a different order)
func anagram(s1, s2 string) bool {
	ls1 := strings.ToLower(s1)
	ls2 := strings.ToLower(s2)
	m1 := makeMap(ls1)
	m2 := makeMap(ls2)

	return reflect.DeepEqual(m1, m2)
}

func makeMap(s string) map[string]int {
	m := make(map[string]int)
	for _, r := range s {
		if _, ok := m[string(r)]; ok {
			m[string(r)]++
		} else {
			m[string(r)] = 1
		}
	}

	return m
}
