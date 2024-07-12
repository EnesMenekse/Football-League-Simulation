package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=1234 dbname=FootballLeague sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	log.Println("Successfully connected to the database!")
}
