// tagoutline outlines the HTML DOM element indented.
package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	log.SetPrefix("tagoutline ")
	log.SetFlags(0)
	for _, url := range os.Args[1:] {
		_, err := fetch(url)
		if err != nil {
			log.Printf("fetch: %v\n", err)
			continue
		}
	}
}

func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return nil, nil
}
