package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	urls := []string{"https://pkg.go.dev", "https://godoc.org", "http://gopl.io", "https://play.golang.org", "https://pkg.go.dev", "https://godoc.org", "https://play.golang.org"}
	function := Func{f: httpGetBody, done: make(chan struct{})}
	m := New(function)
	var n sync.WaitGroup
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(m.fn.done)
	}()

	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
