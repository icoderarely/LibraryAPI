package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/icoderarely/LibraryAPI/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	fmt.Println("DB Connected, everything is working...")

	mux := router.Router()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(" [error] starting server...")
		return
	}
}
