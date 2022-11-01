package main

import (
	"log"
	"net/http"
)

// https://pkg.go.dev/net/http#FileServer
// https://go.dev/doc/articles/wiki/
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.youtube.com/watch?v=jFfo23yIWac&t=611s

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // default points to index.html
	http.Handle("/", fileServer)                        // root route
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/form", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
