package common

import (
	"database/sql"
	"log"
)

// ConnectionString is used for connecting to the local postgres DB
var ConnectionString = `user=postgres host=127.0.0.1 dbname=seu_content_intake sslmode=disable binary_parameters=yes`

// GetDBConnection centralizes the connection calls
func GetDBConnection() *sql.DB {
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		log.Panicf("Error making connection. Stopping app: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Panicf("Error testing connection. Stopping app: %v", err)
	}
	return db
}
