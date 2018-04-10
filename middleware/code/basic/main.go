package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// define simple route
	http.HandleFunc("/hello", helloGET)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//
// Handlers
//

func helloGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
