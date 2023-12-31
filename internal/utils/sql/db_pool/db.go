package db_pool

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	defaultNumConn     = 10
	maxNumConn         = 100
	defaultPoolTimeout = time.Second * 15
	maxPoolTimeout     = time.Minute
)

type DB struct {
	pool chan *Conn
	opt  *Option
}

type Option struct {
	NumConn     int
	PoolTimeout time.Duration
}

func Open(dbPath string, opts ...Option) (*DB, error) {
	var opt Option
	if len(opts) == 0 {
		opt = Option{
			NumConn:     defaultNumConn,
			PoolTimeout: defaultPoolTimeout,
		}
	} else {
		opt = Option{
			NumConn:     min(opts[0].NumConn, maxNumConn),
			PoolTimeout: min(opts[0].PoolTimeout, maxPoolTimeout),
		}
	}

	pool := make(chan *Conn, opt.NumConn)
	g := errgroup.Group{}

	for i := 0; i < opt.NumConn; i++ {
		g.Go(func() error {
			db, err := sql.Open("sqlite3", dbPath)
			if err != nil {
				return err
			}

			select {
			case pool <- &Conn{
				DB:   db,
				pool: pool,
			}:
				return nil
			case <-time.After(opt.PoolTimeout):
				return fmt.Errorf("failed to create connection")
			}
		})
	}

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &DB{
		pool: pool,
		opt:  &opt,
	}, nil
}

func (db *DB) Acquire() (*Conn, error) {
	select {
	case conn := <-db.pool:
		return conn, nil
	case <-time.After(db.opt.PoolTimeout):
		return nil, fmt.Errorf("connections is not available")
	}
}

func (db *DB) Query(ctx context.Context, sql string, args ...interface{}) (*Rows, error) {
	c, err := db.Acquire()
	if err != nil {
		return nil, err
	}
	defer c.Release()

	rows, err := c.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return &Rows{
		Rows: rows,
	}, nil
}

func (db *DB) Exec(ctx context.Context, sql string, args ...interface{}) (*Result, error) {
	c, err := db.Acquire()
	if err != nil {
		return nil, err
	}
	defer c.Release()

	res, err := c.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return &Result{
		Result: res,
	}, nil
}

func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	c, err := db.Acquire()
	if err != nil {
		return nil, err
	}
	defer c.Release()

	tx, err := c.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Tx{
		Tx: tx,
	}, nil
}

func (db *DB) BeginFunc(ctx context.Context, f func(tx *Tx) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			err = rollbackErr
		}
	}()

	fErr := f(tx)
	if fErr != nil {
		_ = tx.Rollback()
		return fErr
	}

	return tx.Commit()
}

func (db *DB) Close() error {
	g := errgroup.Group{}

	for i := 0; i < db.opt.NumConn; i++ {
		g.Go(func() error {
			select {
			case c := <-db.pool:
				return c.Close()
			case <-time.After(db.opt.PoolTimeout):
				return fmt.Errorf("failed to close connection")
			}
		})
	}

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}
