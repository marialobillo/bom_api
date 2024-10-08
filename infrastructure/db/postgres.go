package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq" // Postgres driver
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "user=root password=123456 dbname=bom_api sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database successfully")
}