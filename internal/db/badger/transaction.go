package badger

import (
	"errors"

	"github.com/dgraph-io/badger/v3"

	"github.com/glebziz/fs_db"
)

type transaction struct {
	*badger.Txn
}

func (t transaction) GetAll() ([]Item, error) {
	it := t.Txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	var items []Item
	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		err := item.Value(func(val []byte) error {
			items = append(items, Item{
				Key:   item.Key(),
				Value: val,
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
