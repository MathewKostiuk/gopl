package counters

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

func TestWordCounter(t *testing.T) {
	var w WordCounter
	w.Write([]byte("hello world hello planet"))
	fmt.Println(w)
}

func TestLineCounter(t *testing.T) {
	var l LineCounter
	l.Write([]byte("\n\n\n"))
	fmt.Println(l)
}
