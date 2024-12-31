// fetchlinks fetches and find links, fetch+findlink.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		for _, link := range links {
			fmt.Printf("%s: %s\n", url, link)
		}
	}
}

// findlinks fetch and find links.
func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch(%q): %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code(%q): %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse error(%q): %w", url, err)
	}
	return visit(nil, doc), nil
}

// visit visits the DOM tree and returns the found URL links as a slice.
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
