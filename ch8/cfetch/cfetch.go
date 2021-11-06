package cfetch

import (
	"context"
	"fmt"
	"net/http"
)

// ConcurrentFetch fetches a group of urls and returns the first response
// while cancelling the remaining requests
func ConcurrentFetch(urls []string) *http.Response {
	var ctx context.Context
	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	responses := make(chan *http.Response, len(urls))

	for _, url := range urls {
		go func(url string) {
			res, err := fetch(ctx, url)
			if err != nil {
				fmt.Printf("error in goroutine: %v\n", err)
				return
			}
			responses <- res
		}(url)
	}

	res := <-responses
	cancel()
	return res
}

func fetch(ctx context.Context, url string) (*http.Response, error) {
	if cancelled(ctx) {
		return nil, fmt.Errorf("the request was cancelled")
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error in fetch request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in fetch response: %v", err)
	}
	defer res.Body.Close()
	return res, nil
}

func cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
