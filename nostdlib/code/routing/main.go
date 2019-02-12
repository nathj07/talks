package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
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
	// get the values from the path - inherently buggy as is
	parts := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "Age: %s; Loc: %s", parts[3], parts[4])
}

func namedAgeLocGET(w http.ResponseWriter, r *http.Request,  ps httprouter.Params){
	fmt.Fprintf(w, "Age: %s; Loc: %s", ps.ByName("age"), ps.ByName("loc"))
}

func namedAgeLocPUT(w http.ResponseWriter, r *http.Request,  ps httprouter.Params){
	fmt.Fprintf(w, "PUTTING Age: %s; Loc: %s", ps.ByName("age"), ps.ByName("loc"))
}
func main() {
	// Standard library only
	http.HandleFunc("/", indexHdlr)
	http.HandleFunc("/name/", nameHdlr)           // get one thing from the path
	http.HandleFunc("/age/location/", ageLocHdlr) // get two things from the path
	// err := http.ListenAndServe(":9090", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	// Not very RESTful for those paths.
	// Would rather have /age/:age/location/:location
	// Using httprouter
	router := httprouter.New()
	router.GET("/age/:age/location/:loc", namedAgeLocGET) // using named path components
	router.PUT("/age/:age/location/:loc", namedAgeLocPUT)

	log.Print("Serving on port 9090")
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// NOTE: For TESTING cover testify and httpmock