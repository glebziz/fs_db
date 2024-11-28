package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Delete(_ context.Context, id string) (model.Transaction, error) {
	tx, ok := r.storage.Load(id)
	if !ok {
		return model.Transaction{}, fs_db.ErrTxNotFound
	}

	r.storage.Delete(id)
	return tx, nil
}
