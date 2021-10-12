package htmlparser

import (
	"io"

	"golang.org/x/net/html"
)

type HTMLStringReader struct {
	s string
	i int64
}

func (hr *HTMLStringReader) Read(b []byte) (n int, err error) {
	if hr.i >= int64(len(hr.s)) {
		return 0, io.EOF
	}
	n = copy(b, hr.s[hr.i:])
	hr.i += int64(n)
	return
}

func (hr *HTMLStringReader) NewReader(s string) *HTMLStringReader {
	return &HTMLStringReader{s, 0}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	c := n.FirstChild
	links = visit(links, c)

	return visit(links, n.NextSibling)
}
