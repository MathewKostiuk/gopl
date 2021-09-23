package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MathewKostiuk/countwi"
)

func main() {
	url := os.Args[1:]

	fmt.Println(url)
	words, images, err := countwi.CountWordsAndImages(url[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(words, images)
}
