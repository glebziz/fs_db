package fs_db

import (
	"context"
	"fmt"
	"io"
)

type txCtxFn func(ctx context.Context) context.Context

type tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Tx interface {
	tx
	store
}

type _tx struct {
	tx
	store store
	ctxFn txCtxFn
}

func CreateTx(store store, tx tx, ctxFn txCtxFn) Tx {
	return &_tx{
		tx:    tx,
		store: store,
		ctxFn: ctxFn,
	}
}

func (t *_tx) Set(ctx context.Context, key string, b []byte) error {
	err := t.store.Set(t.ctxFn(ctx), key, b)
	if err != nil {
		return fmt.Errorf("store set: %w", err)
	}

	return nil
}

func (t *_tx) SetReader(ctx context.Context, key string, reader io.Reader, size uint64) error {
	err := t.store.SetReader(t.ctxFn(ctx), key, reader, size)
	if err != nil {
		return fmt.Errorf("store set reader: %w", err)
	}

	return nil
}

func (t *_tx) Get(ctx context.Context, key string) ([]byte, error) {
	b, err := t.store.Get(t.ctxFn(ctx), key)
	if err != nil {
		return nil, fmt.Errorf("store get: %w", err)
	}

	return b, nil
}

func (t *_tx) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	r, err := t.store.GetReader(t.ctxFn(ctx), key)
	if err != nil {
		return nil, fmt.Errorf("store get reader: %w", err)
	}

	return r, nil
}

func (t *_tx) Delete(ctx context.Context, key string) error {
	err := t.store.Delete(t.ctxFn(ctx), key)
	if err != nil {
		return fmt.Errorf("store delete: %w", err)
	}

	return nil
}
