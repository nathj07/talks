package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/nathj07/talks/concurrency/code/const"
)

// Provider is the basic unit we are working with for this example
type Provider struct {
	url  string
	name string
}

var providerChan chan *Provider

func main() {
	// db connection

	db, err := sql.Open("postgres", constants.ConnectionString)
	if err != nil {
		os.Exit(1)
	}

	providerChan = make(chan *Provider) // unbuffered
	// fetch data
	go fetchData(db)
	useData()
}

func fetchData(db *sql.DB) {
	// bounded loop to simulate repeated fetches
	for i := 0; i <= 1; i++ {
		fmt.Println("Iteration ", i)
		fmt.Println("Fetch Data from DB")
		rows, err := db.Query("SELECT name, url FROM provider") // need a way to get distinct sets on each iteration
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
	close(providerChan) // artificial closure for the demo; more to come
}

func useData() {
	for p := range providerChan {
		fmt.Printf("Read from chan: %q, %q\n", p.name, p.url)
	}

}
