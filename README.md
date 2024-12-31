# A web crawler in Go

It demonstrates how easy to write a simple web crawler in Go
programming language, as demonstrated in the book, [The Go Programming
Language][gopl].

[gopl]: https://gopl.io

## Utility commands

- [fetch][] fetches the HTML document for the web site.
- [findlink][] finds URL links from the HTML document.
- [outline][] outlines the HTML DOM elements.
- [fetchlinks][] fetches and show found links, e.g. fetch | findlink
- [waitserver][] waits for the server reachability

Happy hacking!

[fetch]: ./cmd/fetch/main.go
[findlink]: ./cmd/findlink/main.go
[outline]: ./cmd/outline/main.go
[fetchlinks]: ./cmd/fetchlinks/main.go
[waitserver]: ./cmd/waitserver/main.go
