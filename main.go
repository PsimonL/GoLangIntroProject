package main

import (
	"net/http"
)

// https://pkg.go.dev/net/http#FileServer
// https://go.dev/doc/articles/wiki/
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.youtube.com/watch?v=jFfo23yIWac&t=611s

func main() {
	fileServer := http.FileServer()
}
