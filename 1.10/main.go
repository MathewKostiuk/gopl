package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for i, url := range os.Args[1:] {
		go fetch(url, ch, i) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, i int) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	file, err := os.Create("hello.txt")

	if err != nil {
		ch <- fmt.Sprintf("While creating file %s: %v", url, err)
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %7d %s", secs, nbytes, url)
}
