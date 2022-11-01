package main

import (
	"fmt"
	"log"
	"net/http"
)

// https://pkg.go.dev/net/http#FileServer
// https://go.dev/doc/articles/wiki/
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.youtube.com/watch?v=jFfo23yIWac&t=611s

func formHandler(w http.ResponseWriter, r *http.Request) { // response - request
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { // by default used GET
		http.Error(w, "Method from CRUD doesnt supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func helloHandler() {

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // default points to index.html
	http.Handle("/", fileServer)                        // root route
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/form", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // https://www.reddit.com/r/golang/comments/dbw0uc/need_clarification_on_the_use_of_semicolon_after/
	}
}
