package main

import (
	"fmt"

	"github.com/icoderarely/LibraryAPI/internal/db"
)

func main() {
	db := db.ConnectDB()
	defer db.Close()

	fmt.Println("DB Connected, everything is working...")
}
