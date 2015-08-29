package main

import (
	"flag"
	"fmt"
)

var greeting string

// main is the simple entry point taking CL args to launch a conversation. It then starts the
// conversations acordingly
func main() {
	flag.StringVar(&greeting, "greeting", "", "-greeting= to supply your initial greeting.")
	flag.Parse()
	manager(greeting)
}

func manager(greeting string) {
	resp := start(greeting)
	fmt.Println(resp)
}

func start(greeting string) string {

	return greeting
}

// start2 takes a greeting string and at the very least returns that as the response.
// it the greeting is supported then we get a relevant, in language response.
func start2(greeting string) string {

	switch greeting {
	case "Salut":
		return "Salut, ça va ?"
	case "Hola":
		return "Hola, ¿Cómo estás?"
	}
	// return at least the given greeting to seem polite
	return greeting
}

// TODO Show the test of a full workflow through this, need a goodbye and moddle
// conversation for that to be really illustrative.
