package db_pool

import (
	"context"
	"database/sql"
)

type Tx struct {
	*sql.Tx
}

func (tx *Tx) Query(ctx context.Context, sql string, args ...interface{}) (*Rows, error) {
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return &Rows{
		Rows: rows,
	}, nil
}

func (tx *Tx) Exec(ctx context.Context, sql string, arguments ...interface{}) (*Result, error) {
	res, err := tx.ExecContext(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}

	return &Result{
		Result: res,
	}, nil
}

func (tx *Tx) Begin(_ context.Context) (*Tx, error) {
	return tx, nil
}

func (tx *Tx) BeginFunc(_ context.Context, f func(tx *Tx) error) error {
	return f(tx)
}
