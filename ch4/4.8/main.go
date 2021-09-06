// Modify charcount to count letters, digits, and so on in their
// Unicode categories, using functions like unicode.IsLetter
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letters := make(map[rune]int)
	nums := make(map[rune]int)
	spaces := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters[r]++
		}
		if unicode.IsNumber(r) {
			nums[r]++
		}
		if unicode.IsSpace(r) {
			spaces[r]++
		}
		utflen[n]++
	}
	printValues(letters, "letters")
	fmt.Println()
	printValues(nums, "numbers")
	fmt.Println()
	printValues(spaces, "spaces")
	fmt.Println()
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	os.Exit(0)
}

func printValues(vals map[rune]int, name string) {
	fmt.Printf("%s\tcount\n", name)
	for c, n := range vals {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
