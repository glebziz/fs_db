package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

func setupSqlite(ctx context.Context, sqlitePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqlitePath)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}

	goose.SetBaseFS(migrations)
	err = goose.SetDialect("sqlite3")
	if err != nil {
		return nil, fmt.Errorf("set dialect: %w", err)
	}

	err = goose.UpContext(ctx, db, "migrations")
	if err != nil {
		return nil, fmt.Errorf("up migrations: %w", err)
	}

	return db, nil
}
