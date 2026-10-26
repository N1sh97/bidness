package api

import (
	"api/sql"
	"context"
	"fmt"
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

	var count int64
	err = db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM authors",
	).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total authors imported: %d\n", count)
}
