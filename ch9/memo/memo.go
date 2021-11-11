package memo

import (
	"sync"
)

type Func struct {
	f    func(key string) (interface{}, error)
	done chan struct{}
}

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	fn    Func
	mu    sync.Mutex
	cache map[string]*entry
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{fn: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	go func() {
		for {
			<-memo.fn.done
			memo.mu.Lock()
			delete(memo.cache, key)
			memo.mu.Unlock()
			return
		}
	}()
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.fn.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // Wait for ready condition
	}
	return e.res.value, e.res.err
}
