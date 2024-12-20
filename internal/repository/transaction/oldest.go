package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Oldest(_ context.Context) (model.Transaction, error) {
	it := r.storage.Iter()
	if !it.Next() {
		return model.Transaction{}, fs_db.ErrTxNotFound
	}

	return it.Val(), nil
}
