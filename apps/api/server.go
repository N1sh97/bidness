package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RunHTTPServer(setupCtx context.Context, db *sql.DB) error {

	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	setupAuthorsRoutes(router, db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Stashing server on http://localhost%s", srv.Addr)

	go func() {
		<-setupCtx.Done()
		srv.Shutdown(context.Background())
	}()
	return srv.ListenAndServe()

}
