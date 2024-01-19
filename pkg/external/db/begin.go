package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	isoLevel "github.com/glebziz/fs_db/internal/adapter/iso_level"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (db *db) Begin(ctx context.Context, level ...model.TxIsoLevel) (fs_db.Tx, error) {
	var l model.TxIsoLevel
	if len(level) > 0 {
		l = level[0]
	} else {
		l = fs_db.IsoLevelDefault
	}

	resp, err := db.client.BeginTx(ctx, &store.BeginTxRequest{
		IsoLevel: isoLevel.ConvertToGrpc(l),
	})
	if err != nil {
		return nil, fmt.Errorf("begin: %w", grpc.ClientError(err))
	}

	t := tx{
		id:     resp.Id,
		client: db.client,
	}

	return fs_db.CreateTx(db, &t, t.ctx), nil
}
