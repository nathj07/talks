package main

import (
	"database/sql"
	"os"
)

func main() {
	// db connection
	connectionString := `user=postgres host=localhost dbname=seu_content_intake sslmode=disable port=6543 binary_parameters=yes`
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		os.Exit(1)
	}
	BatchFlow(db)
	ContinuousFlow(db)
}

// BatchFlow uses a buffered chan to hold the data to process. This method is the control method for the work.
func BatchFlow(db *sql.DB) {
	// fetch data

	// buffered chan

}

// ContinuousFlow uses an unbuffered chan to continuously write data to for processing. This method is the control method for the work.
func ContinuousFlow(db *sql.DB) {
	// fetch data
	// unbuffered chan
}
