package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres driver
)

type Database struct {
    DB *sql.DB
}

func NewPostgresConnection() (*Database, error) {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, loading environment variables from system")
    }

	user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
        user, password, dbname, host, port, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection is not active: %w", err)
	}

	log.Println("Connected to database successfully")
	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
    return d.DB.Close()
}