// waitserver waits for the server reachability specified by the command line.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetPrefix("waitserver ")
	log.SetFlags(0)
	for _, url := range os.Args[1:] {
		if err := WaitForServer(url); err != nil {
			log.Fatalf("Unreachable: %v\n", err)
		}
	}
}

// WaitForServer wait for the server pointed to by url.
func WaitForServer(url string) error {
	const timeout = 30 * time.Second
	deadline := time.Now().Add(timeout)
	for retries := 0; time.Now().Before(deadline); retries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%v), retrying...", err)
		time.Sleep(time.Second << uint(retries))
	}
	return fmt.Errorf("server failed to respond after %s", timeout)
}
