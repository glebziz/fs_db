package core

import (
	"sync"

	"github.com/glebziz/fs_db/internal/model"
)

const (
	defaultMapCap = 5
)

type Transactions struct {
	m     sync.RWMutex
	store map[string]*Transaction
}

type Transaction struct {
	m     sync.RWMutex
	store map[string]*file
	pool  Pool[file]

	WithoutSearch bool
}

func (txs *Transactions) Get(txId string) (*Transaction, bool) {
	txs.m.RLock()
	defer txs.m.RUnlock()

	if txs.store == nil {
		return nil, false
	}

	tx, ok := txs.store[txId]
	return tx, ok
}

func (txs *Transactions) Put(txId string, tx *Transaction) {
	txs.m.Lock()
	defer txs.m.Unlock()

	if txs.store == nil {
		txs.store = make(map[string]*Transaction, defaultMapCap)
	}

	if tx.store == nil {
		tx.store = make(map[string]*file, defaultMapCap)
	}

	clear(tx.store)
	txs.store[txId] = tx
}

func (txs *Transactions) Delete(txId string) *Transaction {
	txs.m.Lock()
	defer txs.m.Unlock()

	if txs.store == nil {
		return nil
	}

	tx, ok := txs.store[txId]
	if ok {
		delete(txs.store, txId)
	}

	return tx
}

func (tx *Transaction) Lock() {
	if tx == nil {
		return
	}

	tx.m.Lock()
}

func (tx *Transaction) RLock() {
	if tx == nil {
		return
	}

	tx.m.RLock()
}

func (tx *Transaction) Unlock() {
	if tx == nil {
		return
	}

	tx.m.Unlock()
}

func (tx *Transaction) RUnlock() {
	if tx == nil {
		return
	}

	tx.m.RUnlock()
}

func (tx *Transaction) Files() map[string]*file {
	if tx == nil {
		return nil
	}

	return tx.store
}

func (tx *Transaction) File(key string) *file {
	if tx.store == nil {
		return nil
	}

	return tx.store[key]
}

func (tx *Transaction) Len() int {
	if tx == nil {
		return 0
	}

	return len(tx.store)
}

func (tx *Transaction) PushBack(n *Node[model.File]) {
	if tx.store == nil {
		tx.store = make(map[string]*file, defaultMapCap)
	}

	fs, ok := tx.store[n.v.Key]
	if !ok {
		fs = tx.pool.Acquire()
		fs.withoutSearch = tx.WithoutSearch
	}

	fs.PushBack(n)

	if !ok {
		tx.store[n.v.Key] = fs
	}
}

func (tx *Transaction) Clear() {
	if tx == nil {
		return
	}

	fs := make([]*file, 0, defaultMapCap)
	for key, f := range tx.store {
		fs = append(fs, f)
		delete(tx.store, key)
	}

	tx.pool.Release(fs...)
}
