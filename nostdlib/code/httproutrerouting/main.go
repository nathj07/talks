package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)


func ageLocGET(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Age: %s; Loc: %s", ps.ByName("age"), ps.ByName("loc"))
}

func ageLocPUT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PUTTING Age: %s; Loc: %s", ps.ByName("age"), ps.ByName("loc"))
}

func nameAgeLocGET(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Name: %s; Age: %s; Loc: %s",
		ps.ByName("name"), ps.ByName("age"), ps.ByName("loc"))
}

func nameAgeLocPUT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PUTTING Name: %s; Age: %s; Loc: %s",
		ps.ByName("name"),ps.ByName("age"), ps.ByName("loc"))
}

func main() {
	// Std lib only is not very RESTful for those paths.
	// Would rather have /age/:age/location/:location
	// Using httprouter
	router := httprouter.New()
	router.GET("/age/:age/location/:loc", ageLocGET)
	router.PUT("/age/:age/location/:loc", ageLocPUT)
	router.GET("/name/:name/age/:age/location/:loc", nameAgeLocGET)
	router.PUT("/name/:name/age/:age/location/:loc", nameAgeLocPUT)

	log.Print("Serving on port 9090")
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// NOTE: For TESTING cover testify and httpmock

