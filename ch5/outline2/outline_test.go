package outline2

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	b := Outline("https://golang.org")
	r := bytes.NewReader(b)

	_, err := html.Parse(r)
	if err != nil {
		t.Errorf(`Outline("https://golang.org") = false err=%v`, err)
	}
}
