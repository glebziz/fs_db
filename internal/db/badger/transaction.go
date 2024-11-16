package badger

import (
	"errors"
	"slices"

	"github.com/dgraph-io/badger/v3"

	"github.com/glebziz/fs_db"
)

type transaction struct {
	*badger.Txn
}

func (t transaction) GetAll(prefix []byte) ([]Item, error) {
	it := t.Txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	var items []Item
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		item := it.Item()
		err := item.Value(func(val []byte) error {
			items = append(items, Item{
				Key:   slices.Clone(item.Key()),
				Value: slices.Clone(val),
			})

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (t transaction) Get(key []byte) ([]byte, error) {
	item, err := t.Txn.Get(key)
	if errors.Is(err, badger.ErrKeyNotFound) {
		return nil, fs_db.NotFoundErr
	} else if err != nil {
		return nil, err
	}

	var data []byte
	return data, item.Value(func(val []byte) error {
		data = val
		return nil
	})
}
