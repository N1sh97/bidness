package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func setupAuthorsRoutes(r chi.Router, db *sql.DB) {
	r.Route("/authors", func(authorsRouter chi.Router) {
		authorsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {

		})
	})
}
