package main

import "fmt"

func main() {
	// buffered
	cb := make(chan int, 1)
	cb <- 1
	fmt.Println(<-cb)

	// unbuffered
	cu := make(chan int)
	go func() {
		fmt.Println(<-cu)
	}()
	cu <- 2
}
