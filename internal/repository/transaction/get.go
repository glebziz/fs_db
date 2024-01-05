package transaction

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Get(_ context.Context, id string) (*model.Transaction, error) {
	if id == model.MainTxId {
		return &model.Transaction{
			Id:       id,
			IsoLevel: fs_db.IsoLevelDefault,
		}, nil
	}

	tx, ok := r.storage.Load(id)
	if !ok {
		return nil, fs_db.TxNotFoundErr
	}

	return tx, nil
}
