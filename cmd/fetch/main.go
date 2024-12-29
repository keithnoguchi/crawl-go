// fetch fetches the HTML document from URL(s).
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: fetch <URL>...")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Fprintf(
				os.Stderr,
				"Non 200 status code for %q: %s",
				resp.Status,
			)
			os.Exit(1)
		}
		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Read error for %q with %v",
				url, err,
			)
			os.Exit(1)
		}
		fmt.Print(string(data))
	}
}
