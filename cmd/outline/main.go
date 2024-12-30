// outline outlines the HTML DOM elements.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTML parse error: %v", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

// outline outlines the DOM elements.
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := range n.ChildNodes() {
		outline(stack, c)
	}
}
