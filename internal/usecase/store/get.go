package store

import (
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (u *useCase) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	txId := model.GetTxId(ctx)
	tx, err := u.txRepo.Get(ctx, txId)
	if err != nil {
		return nil, fmt.Errorf("tx repository get: %w", err)
	}

	var filter model.FileFilter
	switch tx.IsoLevel {
	case fs_db.IsoLevelReadUncommitted:
	case fs_db.IsoLevelReadCommitted:
		filter.TxId = ptr.Ptr(model.MainTxId)
	case fs_db.IsoLevelRepeatableRead,
		fs_db.IsoLevelSerializable:
		filter.TxId = ptr.Ptr(model.MainTxId)
		filter.BeforeSeq = ptr.Ptr(tx.Seq)
	}

	f, err := u.fRepo.Get(ctx, tx.Id, key, filter)
	if err != nil {
		return nil, fmt.Errorf("file repository get: %w", err)
	}

	cf, err := u.cfRepo.Get(ctx, f.ContentId)
	if err != nil {
		return nil, fmt.Errorf("content file repository get: %w", err)
	}

	content, err := u.cRepo.Get(ctx, cf.Path())
	if err != nil {
		return nil, fmt.Errorf("content repository get: %w", err)
	}

	return content, nil
}
