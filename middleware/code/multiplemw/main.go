// advanced-middleware.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	middlewares := []Middleware{Logging(), Personalise(), Identification()}

	http.HandleFunc("/hello", Chain(helloGET, middlewares))
	http.HandleFunc("/goodbye", Chain(goodbyeGET, middlewares))

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

// Middleware defines the custom type for all our middleware funcs, useful for gathering them together
// START OMIT
type Middleware func(http.HandlerFunc) http.HandlerFunc

// END OMIT
// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
				log.Println("User", r.Context().Value(usernameKey))
			}()
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

type ctxKey int

const (
	usernameKey ctxKey = 1
)

// Identification checks for the  user name in the query string,
// if present it adds it to the context for the logger and prints it on the page.
func Identification() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user := r.FormValue("username")
			if user == "" {
				// nothing more to do here
				f(w, r)
			}
			ctx = context.WithValue(ctx, usernameKey, user)
			r = r.WithContext(ctx)
			f(w, r)
		}
	}
}

// Personalise uses the context and prints the username if found
func Personalise() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if username := r.Context().Value(usernameKey); username != nil {
				w.Write([]byte(fmt.Sprintf("User: %s\n", username)))
			}
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, mw []Middleware) http.HandlerFunc {
	for _, m := range mw {
		f = m(f)
	}
	return f
}
