package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/icoderarely/LibraryAPI/internal/db"
	"github.com/icoderarely/LibraryAPI/internal/models"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error xyz", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := db.CreateBook(&book); err != nil {
		http.Error(w, "Error creating book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := map[string]interface{}{
		"status": "success",
		"data":   book,
	}

	json.NewEncoder(w).Encode(resp)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	book, err := db.GetBook(id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := db.GetBooks()
	if err != nil {
		http.Error(w, fmt.Sprintf("%e", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
