package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to your first web server!")
}

func nameGet(w http.ResponseWriter, r *http.Request) {
	// name is the last part of the path
	fmt.Fprintf(w, "From the path: %q", strings.Split(r.URL.Path, "/")[2])

}

func ageLocGet(w http.ResponseWriter, r *http.Request) {
	// name is the last part of the path
	parts := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "Age: %s; Loc: %s", parts[3], parts[4])

}

func main() {
	http.HandleFunc("/", indexGet) // set router
	http.HandleFunc("/name/", nameGet) // get one thing from the path
	http.HandleFunc("/age/location/", ageLocGet) // get two things from the path
	// Not very RESTful for those paths.
	// Would rather have /age/:age/location/:location
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}