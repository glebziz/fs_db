package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Delete(_ context.Context, id string) (*model.Transaction, error) {
	tx, ok := r.storage.Load(id)
	if !ok {
		return nil, fs_db.TxNotFoundErr
	}

	r.storage.Delete(id)
	return tx, nil
}
