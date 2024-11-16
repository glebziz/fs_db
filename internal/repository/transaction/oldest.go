package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Oldest(_ context.Context) (model.Transaction, error) {
	it := r.storage.Iter()
	if !it.Next() {
		return model.Transaction{}, fs_db.TxNotFoundErr
	}

	return it.Val(), nil
}
