package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	_"time"

	"github.com/go-chi/chi/v5"
)

type AuthorData struct {
	Key            string    `json:"key"`
	Name           string    `json:"name"`
	TypeKey        string    `json:"type_key"`
}

func setupAuthorsRoutes(r chi.Router, db *sql.DB) {
	r.Route("/authors", func(authorsRouter chi.Router) {
		authorsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			getAuthors(w, r, db)
		})
	})
}

func getAuthors(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.QueryContext(r.Context(), `
		SELECT
			key,
			name,
			type_key
		FROM authors
	`)
	if err != nil {
		log.Printf("query authors failed: %v", err)
		http.Error(w, "failed to get authors", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	authors := make([]AuthorData, 0)

	for rows.Next() {
		var a AuthorData

		if err := rows.Scan(
			&a.Key,
			&a.Name,
			&a.TypeKey,
		); err != nil {
			log.Printf("failed to scan author row: %v", err)
			http.Error(w, "failed to read author", http.StatusInternalServerError)
			return
		}

		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "failed to iterate authors", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		return
	}
}
