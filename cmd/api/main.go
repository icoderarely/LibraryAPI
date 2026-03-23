package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/icoderarely/LibraryAPI/internal/db"
	"github.com/icoderarely/LibraryAPI/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db := db.ConnectDB()
	defer db.Close()

	fmt.Println("DB Connected, everything is working...")

	router := router.Router()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(" [error] starting server...")
		return
	}
}
