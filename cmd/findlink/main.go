// findlink finds the link from the HTML DOM and print out the links
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	docs, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTML document parse error: %v", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, docs) {
		fmt.Println(link)
	}
}

// visit visits DOM nodes in breadth-first search and returns list of URLs.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := range n.ChildNodes() {
		links = visit(links, c)
	}
	return links
}
