package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexHdlr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to your first web server!")
}

func nameHdlr(w http.ResponseWriter, r *http.Request) {
	// name is the last part of the path
	fmt.Fprintf(w, "From the path: %q", strings.Split(r.URL.Path, "/")[2])

}

func ageLocHdlr(w http.ResponseWriter, r *http.Request) {
	// test verb
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Only GET requests are support currently, you made a %s request", r.Method)
		return
	}
	// get the values from the path
	parts := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "Age: %s; Loc: %s", parts[3], parts[4])
}

func ageLocNameHdlr(w http.ResponseWriter, r *http.Request) {
	// test verb
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Only GET requests are support currently, you made a %s request", r.Method)
		return
	}
	// get the values from the path
	parts := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "Age: %s; Loc: %s; Name: %s", parts[3], parts[4], parts[5])
}

func main() {
	// Standard library only
	http.HandleFunc("/", indexHdlr)
	http.HandleFunc("/name/", nameHdlr)           // get one thing from the path
	http.HandleFunc("/age/location/", ageLocHdlr) // get two things from the path
	// this, with no vars, would clash with the above, and panic as there are not enough path parts
	http.HandleFunc("/age/location/name", ageLocNameHdlr)
	log.Print("Serving on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// NOTE: For TESTING cover testify and httpmock
