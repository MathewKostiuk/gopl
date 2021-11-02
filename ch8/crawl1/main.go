package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enfore a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var d = flag.Int("depth", 1, "depth limits the crawler to the specified value")

type wl struct {
	list  []string
	depth int
}

func main() {
	flag.Parse()
	worklist := make(chan wl)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments
	n++
	go func() { worklist <- wl{flag.Args()[:], *d} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth < 0 {
			continue
		}
		fmt.Println(list.depth)
		for _, link := range list.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link, list.depth)
				}(link)
			}
		}
	}
}

func crawl(url string, dep int) wl {
	dep--
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return wl{list, dep}
}
