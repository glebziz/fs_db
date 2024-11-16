package core

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Get(_ context.Context, txId, key string, filter model.FileFilter) (model.File, error) {
	var f, s model.File
	if filter.BeforeSeq == nil && filter.TxId == nil {
		u.allStore.RLock()
		defer u.allStore.RUnlock()

		f = u.allStore.File(key).Latest()
	} else if filter.TxId != nil {
		tx, ok := u.txStore.Get(txId)
		if ok {
			f = tx.File(key).Latest()
		}
		tx.RLock()
		defer tx.RUnlock()

		sTx, ok := u.txStore.Get(*filter.TxId)
		if ok {
			sTx.RLock()
			defer sTx.RUnlock()

			if filter.BeforeSeq == nil {
				s = sTx.File(key).Latest()
			} else {
				s = sTx.File(key).LastBefore(*filter.BeforeSeq)
			}
		}
	}

	latest := f.Latest(s)
	if latest.Seq.Zero() || latest.Deleted() {
		return model.File{}, fs_db.NotFoundErr
	}

	return latest, nil
}
