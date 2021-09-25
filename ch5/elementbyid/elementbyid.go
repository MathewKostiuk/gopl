package elementbyid

import (
	"golang.org/x/net/html"
)

func ElementById(doc *html.Node, id string) (*html.Node, bool) {
	return forEachNode(doc, id, startElement, nil)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) (*html.Node, bool) {
	var found bool

	if pre != nil && !found {
		found = pre(n, id)
	}

	if found {
		return n, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node, f := forEachNode(c, id, pre, post)
		if f {
			return node, true
		}
	}

	if post != nil && !found {
		post(n, id)
	}

	return nil, false
}

func startElement(n *html.Node, id string) bool {
	var found bool
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				found = true
				break
			}
		}
	}
	return found
}
