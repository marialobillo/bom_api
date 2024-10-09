package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq" // Postgres driver
)

type Database struct {
    DB *sql.DB
}

func NewPostgresConnection() *Database {
	//connStr := "user=postgres password=postgres dbname=bom_api sslmode=disable"
	connStr := "user=root password=123456 dbname=bom_api sslmode=disable"
	//connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.Ping(); err != nil {
        log.Fatal("Database connection is not active:", err)
    }
	log.Println("Connected to database successfully")

	return &Database{DB: db}
}

func (d *Database) Close() error {
    return d.DB.Close()
}