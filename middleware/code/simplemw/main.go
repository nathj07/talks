package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", logging(hello))
	http.HandleFunc("/goodbye", logging(goodbye))

	http.ListenAndServe(":8080", nil)
}

//
// Handlers
//

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "goodbye")
}

//
// Middleware
//

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		defer log.Println("Completed Path")
		next.ServeHTTP(w, r)
	}
}
