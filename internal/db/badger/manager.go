package badger

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v3"

	"github.com/glebziz/fs_db/internal/model/transactor"
)

type Item struct {
	Key   []byte
	Value []byte
}

type Manager struct {
	db *badger.DB
}

type QueryManager interface {
	Set(key []byte, val []byte) error
	GetAll(prefix []byte) ([]Item, error)
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
}

type Provider interface {
	transactor.Transactor
	DB(ctx context.Context) QueryManager
}

type ctxTxn struct{}

func New(dbPath string) (*Manager, error) {
	db, err := badger.Open(badger.DefaultOptions(dbPath).WithLogger(nil))
	if err != nil {
		return nil, fmt.Errorf("badger open: %w", err)
	}

	return &Manager{
		db: db,
	}, nil
}

func (m *Manager) Set(key []byte, val []byte) error {
	return m.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (m *Manager) GetAll(prefix []byte) (items []Item, err error) {
	return items, m.db.View(func(txn *badger.Txn) error {
		items, err = transaction{txn}.GetAll(prefix)
		if err != nil {
			return err
		}

		return nil
	})
}

func (m *Manager) Get(key []byte) (data []byte, err error) {
	return data, m.db.View(func(txn *badger.Txn) error {
		data, err = transaction{txn}.Get(key)
		if err != nil {
			return err
		}

		return nil
	})
}

func (m *Manager) Delete(key []byte) error {
	return m.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (m *Manager) GC() {
	m.db.RunValueLogGC(0.5) //nolint:errcheck,mnd
}

func (m *Manager) DB(ctx context.Context) QueryManager {
	txn, ok := ctx.Value(ctxTxn{}).(*badger.Txn)
	if ok && txn != nil {
		return transaction{txn}
	}

	return m
}

func (m *Manager) RunTransaction(ctx context.Context, fn transactor.TransactionFn) error {
	querier, ok := ctx.Value(ctxTxn{}).(QueryManager)
	if ok && querier != nil {
		return fn(ctx)
	}

	return m.db.Update(func(txn *badger.Txn) error {
		return fn(context.WithValue(ctx, ctxTxn{}, txn))
	})
}

func (m *Manager) Close() error {
	return m.db.Close()
}
