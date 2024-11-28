package store

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (u *UseCase) GetKeys(ctx context.Context) ([]string, error) {
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

	files, err := u.fRepo.GetFiles(ctx, tx.Id, filter)
	if err != nil {
		return nil, fmt.Errorf("file repository get files: %w", err)
	}

	keys := make([]string, 0, len(files))
	for _, file := range files {
		_, err = u.cfRepo.Get(ctx, file.ContentId)
		if errors.Is(err, fs_db.ErrNotFound) {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("content file repository get: %w", err)
		}

		keys = append(keys, file.Key)
	}

	sort.Strings(keys)

	return keys, nil
}
