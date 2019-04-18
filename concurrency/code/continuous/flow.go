package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
	"github.com/nathj07/talks/concurrency/code/common"
)

// Provider is the basic unit we are working with for this example
type Provider struct {
	url  string
	name string
}

var providerChan chan *Provider
var concRate int

func main() {
	flag.IntVar(&concRate, "concrate", 1, "set the worker pool size, default to 1 worker")
	flag.Parse()

	db := common.GetDBConnection()
	providerChan = make(chan *Provider) // unbuffered
	// fetch data
	go fetchData(db)
	// blocking calls
	useData(concRate) // using a worker pool to interleave the tasks
}

func fetchData(db *sql.DB) {
	// bounded loop to simulate repeated fetches
	for i := 0; i < 2; i++ {
		fmt.Println("Iteration ", i)
		fmt.Println("Fetch Data from DB")
		rows, err := db.Query(`SELECT publisher_name, url FROM publication_stats LIMIT 10`)
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
			fmt.Printf("Write to chan: %q, %q\n", p.name, p.url)
			providerChan <- p
		}
	}
	close(providerChan) // artificial closure for the demo
}

// useData will read until the chan is closed but it will read items concurrently and the
// work is controlled with a loop acting as a pool and a WaitGroup to ensure it all finishes before we return.
func useData(concRate int) {
	var wg sync.WaitGroup
	// concurrencyRate := 5 // in the wild you'd use a config variable for this
	for i := 0; i < concRate; i++ {
		fmt.Println("Worker ", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range providerChan {
				func() {
					fmt.Printf("Read from chan: %q, %q\n", p.name, p.url)
					resp, err := http.Head(p.url)
					if err != nil {
						fmt.Printf("Error making head request for %q: %v\n", p.url, err)
						return
					}
					defer resp.Body.Close()
					fmt.Printf("Processing Data: %q\t%v\n", p.name, resp.Status)
				}()
			}
		}()
	}
	wg.Wait()
}
