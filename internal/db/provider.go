package db

import (
	"context"
	"database/sql"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func Provide(dbPath string) (*Queries, error) {
	// TODO: configurable sqlite file
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(context.Background(), ddl); err != nil {

		return nil, err
	}

	return New(db), nil
}
