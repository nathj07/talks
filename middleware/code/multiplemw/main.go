// advanced-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	middlewares := []Middleware{Method("GET"), Logging()}

	http.HandleFunc("/hello", Chain(hello, middlewares...))
	http.HandleFunc("/goodbye", Chain(goodbye, middlewares...))

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

// Middleware defines the custom type for all our middleware funcs, useful for gathering them together
// START OMIT
type Middleware func(http.HandlerFunc) http.HandlerFunc
// END OMIT
// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				log.Printf("Method: %s unsupported for path %q\n", r.Method, r.URL.String())
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}


