package main

import (
	"fmt"

	parser "github.com/lukemoran01/htmlparse"
	"golang.org/x/net/html"
)

func main() {
	htmlToParse := parser.HtmlToReader("html/ex3.html")
	htmlTree, _ := html.Parse(htmlToParse)

	fmt.Println(parser.Dfs(htmlTree))
}
