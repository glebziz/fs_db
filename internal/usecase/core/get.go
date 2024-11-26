package core

import (
	"context"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *UseCase) Get(_ context.Context, txId, key string, filter model.FileFilter) (model.File, error) {
	var f, s model.File
	if filter.BeforeSeq == nil && filter.TxId == nil {
		f = u.getFileFromTx(&u.allStore, key, nil)
	} else if filter.TxId != nil {
		tx, ok := u.txStore.Get(txId)
		if ok {
			f = u.getFileFromTx(tx, key, nil)
		}

		sTx, ok := u.txStore.Get(*filter.TxId)
		if ok {
			s = u.getFileFromTx(sTx, key, filter.BeforeSeq)
		}
	}

	latest := f.Latest(s)
	if latest.Seq.Zero() {
		return model.File{}, fs_db.NotFoundErr
	}

	return latest, nil
}

func (u *UseCase) getFileFromTx(tx *core.Transaction, key string, beforeSeq *sequence.Seq) model.File {
	tx.RLock()
	defer tx.RUnlock()

	var f model.File
	if beforeSeq == nil {
		f = tx.File(key).Latest()
	} else {
		f = tx.File(key).LastBefore(*beforeSeq)
	}

	return f
}
