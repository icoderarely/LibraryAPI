package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/icoderarely/LibraryAPI/internal/models"
)

var ErrNotFound = errors.New("user not found")

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

func GetBook(id int) (models.Book, error) {
	db := ConnectDB()
	defer db.Close()

	query := "SELECT id, title, author, genre, published_year, available, created_at FROM books WHERE id = ?"

	var book models.Book

	err := db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.PublishedYear, &book.Available, &book.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Book{}, ErrNotFound
	}

	return book, nil
}

func GetBooks() ([]models.Book, error) {
	db := ConnectDB()
	defer db.Close()

	query := "SELECT id, title, author, genre, published_year, available, created_at FROM books"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Genre,
			&book.PublishedYear,
			&book.Available,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
