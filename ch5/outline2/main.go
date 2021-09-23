package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			err = fmt.Errorf("error in response: %s", err)
			log.Fatal(err)
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("error parsing: %s", err)
			log.Fatal(err)
		}

		forEachNode(doc, startElement, endElement)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
			}
			fmt.Printf("/>\n")
		}

		if n.FirstChild != nil {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
			}
			fmt.Printf(">\n")
			depth++
		}
	case html.CommentNode:
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	case html.TextNode:
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
