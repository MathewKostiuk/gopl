package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// tokens is a counting semaphore used to
// enfore a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var d = flag.Int("depth", 1, "depth limits the crawler to the specified value")

var ctx context.Context

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

	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read 1 byte.
		cancel()
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		select {
		case <-ctx.Done():
			close(worklist)
			for range worklist {
				// do nothing.
			}
			panic("process cancelled")
		case list := <-worklist:
			if list.depth < 0 {
				continue
			}
			for _, link := range list.list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						select {
						case <-ctx.Done():
							return
						case worklist <- crawl(link, list.depth, ctx, cancel):
						}

					}(link)
				}
			}
		}
	}
}

func crawl(url string, dep int, ctx context.Context, cancel context.CancelFunc) wl {
	dep--
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := extract(url, ctx, cancel)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return wl{list, dep}
}

func extract(url string, ctx context.Context, cancel context.CancelFunc) ([]string, error) {
	if cancelled() {
		return nil, fmt.Errorf("cancelled")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func cancelled() bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
