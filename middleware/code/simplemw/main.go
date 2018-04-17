package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", logging(helloGET))
	http.HandleFunc("/goodbye", logging(goodbyeGET))

	log.Println("App started, listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

//
// Handlers
//

func helloGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func goodbyeGET(w http.ResponseWriter, r *http.Request) {
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
