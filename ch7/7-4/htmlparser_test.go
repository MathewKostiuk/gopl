package htmlparser

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	var hr HTMLStringReader
	r := hr.NewReader(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta http-equiv="X-UA-Compatible" content="IE=edge">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Document</title>
	</head>
	<body>
	  <a href="google.com" class="href"></a>
	  <a href="facebook.com" class="href"></a>
	</body>
	</html>
	`)
	doc, _ := html.Parse(r)

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
