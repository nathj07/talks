package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nathj07/talks/concurrency/code/common"
)

// Provider is the basic unit we are working with for this example
type Provider struct {
	url  string
	name string
}

var providerChan chan *Provider
var done chan struct{}

func main() {
	db := common.GetDBConnection()
	providerChan = make(chan *Provider) // unbuffered data chan
	done = make(chan struct{})          // unbuffered control chans
	// fetch data
	go fetchData(db)
	go useData()
	// blocking call
	stop()
}

func stop() {
	fmt.Println("Blocked in stop")
	<-done // indicates we have drained the channel and can safely stop
	fmt.Println("Stopping")
	return
}

func fetchData(db *sql.DB) {
	defer close(providerChan)
	for i := 0; i <= 1; i++ {
		fmt.Println("Iteration ", i)
		fmt.Println("Fetch Data from DB")
		rows, err := db.Query("SELECT name, url FROM provider")
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			p := &Provider{}
			if err := rows.Scan(&p.name, &p.url); err != nil {
				fmt.Printf("Error in scan: %v", err)
				continue
			}
			fmt.Printf("Write to chan: %q\n", p.name)
			providerChan <- p
		}
	}
}

// useData will read until the chan is closed but it will read items concurrently and the
// work is controlled with a loop acting as a pool and a WaitGroup to ensure it all finishes before we return.
func useData() {
	for p := range providerChan {
		func() {
			fmt.Printf("Read from chan: %q, %q\n", p.name, p.url)
			resp, err := http.Head(p.url)
			if err != nil {
				fmt.Printf("Error making head request for %q: %v\n", p.url, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Processing Data: %q\t%v\n", p.name, resp)
		}()
	}
}
