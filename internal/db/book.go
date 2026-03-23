package db

import (
	"errors"
	"time"

	"github.com/icoderarely/LibraryAPI/internal/models"
)

func CreateBook(book *models.Book) error {
	db := ConnectDB()
	defer db.Close()

	query := "INSERT INTO books (title, author, genre, published_year, available, created_at) VALUES (?, ?, ?, ?, ?, ?)"

	book.CreatedAt = time.Now().UTC()
	result, err := db.Exec(query, book.Title, book.Author, book.Genre, book.PublishedYear, book.Available, book.CreatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)

	return nil
}
