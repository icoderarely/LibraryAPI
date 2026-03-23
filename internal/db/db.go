package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DSN")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("cannot reach db:", err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	return db
}
