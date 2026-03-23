package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/icoderarely/LibraryAPI/internal/db"
	"github.com/icoderarely/LibraryAPI/internal/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error xyz", http.StatusBadRequest)
		return
	}

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
