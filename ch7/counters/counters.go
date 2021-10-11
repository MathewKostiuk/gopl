package counters

import (
	"bufio"
	"strings"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	c := count(p, bufio.ScanWords)
	*w += WordCounter(c)

	return c, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	c := count(p, bufio.ScanLines)
	*l += LineCounter(c)

	return c, nil
}

func count(p []byte, fn bufio.SplitFunc) int {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)
	c := 0
	for scanner.Scan() {
		c++
	}

	return c
}
