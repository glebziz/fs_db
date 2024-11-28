package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Get(_ context.Context, id string) (model.Transaction, error) {
	if id == model.MainTxId {
		return model.Transaction{
			Id:       id,
			IsoLevel: fs_db.IsoLevelDefault,
		}, nil
	}

	tx, ok := r.storage.Load(id)
	if !ok {
		return model.Transaction{}, fs_db.ErrTxNotFound
	}

	return tx, nil
}
