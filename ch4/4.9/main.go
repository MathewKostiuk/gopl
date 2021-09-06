// Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before
// the first class to Scan to break the input into words instead of lines
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	file, err := os.Open("sample_text.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	fmt.Println("Word\tFrequency")
	for w, f := range counts {
		fmt.Printf("%q\t%d\n", w, f)
	}
}
