//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//// https://pkg.go.dev/net/http#FileServer
//// https://go.dev/doc/articles/wiki/
//// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
//// https://www.youtube.com/watch?v=jFfo23yIWac&t=611s
//
//func formHandler(w http.ResponseWriter, r *http.Request) {
//	if err := r.ParseForm(); err != nil {
//		fmt.Fprintf(w, "ParseForm() err: %v", err)
//		return
//	}
//	fmt.Fprintf(w, "POST ok")
//	name := r.FormValue("name")
//	surname := r.FormValue("surname")
//	fmt.Fprintf(w, "name = %s\n", name)
//	fmt.Fprintf(w, "surname = %s\n", surname)
//}
//
//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/hello" {
//		http.Error(w, "404 not found", http.StatusNotFound)
//		return
//	}
//	if r.Method != "GET" { // by default used GET
//		http.Error(w, "method is not supported", http.StatusNotFound)
//		return
//	}
//	fmt.Fprintf(w, "hello!")
//}
//
//func main() {
//	fileServer := http.FileServer(http.Dir("./static")) // default points to index.html
//	http.Handle("/", fileServer)                        // root route
//	http.HandleFunc("/form", formHandler)
//	http.HandleFunc("/hello", helloHandler)
//	fmt.Printf("Starting server at port 8080\n")
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		log.Fatal(err) // https://www.reddit.com/r/golang/comments/dbw0uc/need_clarification_on_the_use_of_semicolon_after/
//	}
//}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	// 1st route - home page
	// 2nd route - "Hello World"
	// 3rd route - form
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
