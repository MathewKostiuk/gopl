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

	pc(doc)
}

func pc(n *html.Node) {
	if n == nil ||
		n.Data == "script" ||
		n.Data == "style" {
		return
	}

	if n.Type == html.TextNode {
		fmt.Printf("%s\n", n.Data)
	}

	pc(n.FirstChild)
	pc(n.NextSibling)
}
