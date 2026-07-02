package db

import (
	"alex-laycalvert/t/internal/config"
	"alex-laycalvert/t/internal/utils"
	"context"
	"database/sql"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func Provide(cfg *config.Config) (*Queries, error) {
	dbPath := utils.ExpandHomeDir(cfg.Get(config.DBPathKey))
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(context.Background(), ddl); err != nil {

		return nil, err
	}

	return New(db), nil
}
