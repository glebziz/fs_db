package sqlite

import (
	"context"
	"embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"

	"github.com/glebziz/fs_db/internal/utils/sql/db_pool"
)

//go:embed migrations/*.sql
var migrations embed.FS

type Manager struct {
	db *db_pool.DB
}

type QueryManager interface {
	Query(ctx context.Context, sql string, args ...interface{}) (*db_pool.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (*db_pool.Result, error)
}

type Provider interface {
	DB(ctx context.Context) QueryManager
}

const ctxQuerier = "ctxQuerierKey"

func New(ctx context.Context, dbPath string) (*Manager, error) {
	db, err := db_pool.Open(dbPath)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	c, err := db.Acquire()
	if err != nil {
		return nil, fmt.Errorf("acquire: %w", err)
	}
	defer c.Release()

	goose.SetBaseFS(migrations)
	err = goose.SetDialect("sqlite3")
	if err != nil {
		return nil, fmt.Errorf("set dialect: %w", err)
	}

	err = goose.UpContext(ctx, c.DB, "migrations")
	if err != nil {
		return nil, fmt.Errorf("up migrations: %w", err)
	}

	return &Manager{db}, nil
}

func (m *Manager) DB(ctx context.Context) QueryManager {
	querier, ok := ctx.Value(ctxQuerier).(QueryManager)
	if ok && querier != nil {
		return querier
	}

	return m.db
}

func (m *Manager) Close() error {
	return m.db.Close()
}
