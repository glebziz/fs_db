package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (db *db) Begin(ctx context.Context, level ...model.TxIsoLevel) (fs_db.Tx, error) {
	var l model.TxIsoLevel
	if len(level) > 0 {
		l = level[0]
	} else {
		l = fs_db.IsoLevelDefault
	}

	txId, err := db.txUc.Begin(ctx, l)
	if err != nil {
		return nil, fmt.Errorf("tx usecase begin: %w", err)
	}

	t := tx{
		id:   txId,
		txUc: db.txUc,
	}

	return fs_db.CreateTx(db, &t, t.ctx), nil
}
