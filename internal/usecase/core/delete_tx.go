package core

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) DeleteTx(_ context.Context, txId string) []model.File {
	tx := u.txStore.Delete(txId)
	if tx == nil {
		return nil
	}

	tx.Lock()
	defer func() {
		tx.Unlock()
		u.txPool.Release(tx)
	}()

	u.allStore.Lock()
	defer u.allStore.Unlock()

	deleteFiles := make([]model.File, 0, tx.Len())
	for _, f := range tx.Files() {
		for n := f.PopFront(); n != nil; n = f.PopFront() {
			deleteFiles = append(deleteFiles, n.V())
			link := n.DeleteLink()
			u.nodePool.Release(link, n)
		}
	}

	return deleteFiles
}
