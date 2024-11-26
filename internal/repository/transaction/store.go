package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Store(_ context.Context, tx model.Transaction) error {
	_, ok := r.storage.Load(tx.Id)
	if ok {
		return fs_db.TxAlreadyExistsErr
	}

	r.storage.Store(tx.Id, tx)

	return nil
}
