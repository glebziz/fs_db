package core

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *UseCase) DeleteOld(_ context.Context, txId string, beforeSeq sequence.Seq) []model.File {
	tx, ok := u.txStore.Get(txId)
	if !ok {
		return nil
	}

	tx.Lock()
	u.allStore.Lock()

	deleteFiles := make([]model.File, 0, tx.Len())
	for _, f := range tx.Files() {
		for file := range f.IterateBeforeSeq(beforeSeq) {
			deleteFiles = append(deleteFiles, file)

			n := f.PopFront()
			link := n.DeleteLink()

			u.nodePool.Release(link, n)
		}
	}

	u.allStore.Unlock()
	tx.Unlock()

	return deleteFiles
}
