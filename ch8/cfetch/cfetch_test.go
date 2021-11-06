package cfetch

import (
	"fmt"
	"testing"
)

func TestConcurrentFetch(t *testing.T) {
	urls := []string{"http://gopl.io", "http://google.com", "http://facebook.com"}
	res := ConcurrentFetch(urls)
	fmt.Println(res)
}
