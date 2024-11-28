package fs_db

import (
	"context"
	"fmt"
	"io"
)

type txCtxFn func(ctx context.Context) context.Context

// TxOps provides transactional operations.
type TxOps interface {
	// Commit commits the transaction.
	Commit(ctx context.Context) error

	// Rollback rolls back the transaction.
	Rollback(ctx context.Context) error
}

// Tx provides fs db transaction interface.
type Tx interface {
	TxOps
	Store
}

type tx struct {
	TxOps
	store Store
	ctxFn txCtxFn
}

// CreateTx returns transaction fs db.
func CreateTx(store Store, t TxOps, ctxFn txCtxFn) Tx {
	return &tx{
		TxOps: t,
		store: store,
		ctxFn: ctxFn,
	}
}

func (t *tx) Set(ctx context.Context, key string, b []byte) error {
	err := t.store.Set(t.ctxFn(ctx), key, b)
	if err != nil {
		return fmt.Errorf("store set: %w", err)
	}

	return nil
}

func (t *tx) SetReader(ctx context.Context, key string, reader io.Reader) error {
	err := t.store.SetReader(t.ctxFn(ctx), key, reader)
	if err != nil {
		return fmt.Errorf("store set reader: %w", err)
	}

	return nil
}

func (t *tx) Get(ctx context.Context, key string) ([]byte, error) {
	b, err := t.store.Get(t.ctxFn(ctx), key)
	if err != nil {
		return nil, fmt.Errorf("store get: %w", err)
	}

	return b, nil
}

func (t *tx) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	r, err := t.store.GetReader(t.ctxFn(ctx), key)
	if err != nil {
		return nil, fmt.Errorf("store get reader: %w", err)
	}

	return r, nil
}

func (t *tx) GetKeys(ctx context.Context) ([]string, error) {
	keys, err := t.store.GetKeys(t.ctxFn(ctx))
	if err != nil {
		return nil, fmt.Errorf("store get keys: %w", err)
	}

	return keys, nil
}

func (t *tx) Delete(ctx context.Context, key string) error {
	err := t.store.Delete(t.ctxFn(ctx), key)
	if err != nil {
		return fmt.Errorf("store delete: %w", err)
	}

	return nil
}

func (t *tx) Create(ctx context.Context, key string) (File, error) {
	wc, err := t.store.Create(t.ctxFn(ctx), key)
	if err != nil {
		return nil, fmt.Errorf("store create: %w", err)
	}

	return wc, nil
}
