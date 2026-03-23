package router

import (
	"net/http"

	"github.com/icoderarely/LibraryAPI/internal/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	mux.HandleFunc("POST /books", handlers.CreateBook)

	return mux
}
