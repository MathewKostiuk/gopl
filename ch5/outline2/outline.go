package outline2

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var depth int

func Outline(url string) []byte {
	var b bytes.Buffer

	fmt.Fprintf(&b, "<!DOCTYPE html>\n")

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

	forEachNode(&b, doc, startElement, endElement)
	return b.Bytes()
}

func forEachNode(b *bytes.Buffer, n *html.Node, pre, post func(b *bytes.Buffer, n *html.Node)) {
	if pre != nil {
		pre(b, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(b, c, pre, post)
	}

	if post != nil {
		post(b, n)
	}
}

func startElement(b *bytes.Buffer, n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.FirstChild == nil {
			fmt.Fprintf(b, "%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Fprintf(b, " %s=\"%s\"", a.Key, a.Val)
			}
			if n.Data != "script" {
				fmt.Fprintf(b, "/>\n")
			} else {
				fmt.Fprintf(b, ">")
			}
		}

		if n.FirstChild != nil {
			fmt.Fprintf(b, "%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Fprintf(b, " %s=\"%s\"", a.Key, a.Val)
			}
			fmt.Fprintf(b, ">\n")
			depth++
		}
	case html.CommentNode:
		fmt.Fprintf(b, "%*s%s\n", depth*2, "", n.Data)
	case html.TextNode:
		fmt.Fprintf(b, "%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(b *bytes.Buffer, n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Fprintf(b, "%*s</%s>\n", depth*2, "", n.Data)
	}
	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data == "script" {
		fmt.Fprintf(b, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
