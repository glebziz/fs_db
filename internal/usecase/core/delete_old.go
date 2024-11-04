package core

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) DeleteOld(_ context.Context, txId string, beforeSeq sequence.Seq) {
	tx, ok := u.txStore.Get(txId)
	if !ok {
		return
	}

	tx.Lock()
	u.allStore.Lock()

	deleteFiles := make([]model.File, 0, tx.Len())
	defer func() {
		u.m.Lock()
		u.deleteFiles = append(u.deleteFiles, deleteFiles...)
		u.m.Unlock()
	}()
	for _, f := range tx.Files() {
		nextFn := f.IterateBeforeSeq(beforeSeq)
		for n := nextFn(); n != nil; n = nextFn() {
			deleteFiles = append(deleteFiles, n.V())
			n.DeleteLink() // TODO free link node
			n.Delete()     // TODO free node

		}
	}

	u.allStore.Unlock()
	tx.Unlock()
}
