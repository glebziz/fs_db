package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Store(_ context.Context, tx model.Transaction) error {
	_, ok := r.storage.LoadOrStore(tx.Id, &tx)
	if ok {
		return fs_db.TxAlreadyExistsErr
	}

	return nil
}
