package main

import (
	"flag"
	"fmt"

	"github.com/nathj07/talks/unit_test/code/conversation"
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
	resp := conversation.Greeting(greeting)
	fmt.Println(resp)
}
