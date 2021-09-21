package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	dm := make(map[string]int)

	m := mapEls(dm, doc)
	for el, c := range m {
		fmt.Printf("%s: %d\n", el, c)
	}
}

func mapEls(dm map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return dm
	}
	if n.Type == html.ElementNode {
		dm[n.Data]++
	}

	c := n.FirstChild
	dm = mapEls(dm, c)

	return mapEls(dm, n.NextSibling)
}
