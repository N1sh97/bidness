package api

import (
	"api/sql"
	"context"
	"log"
)

func Main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.SetupDB(ctx, "./data", false)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := RunHTTPServer(ctx, db); err != nil {
		log.Fatal(err)
	}
}
