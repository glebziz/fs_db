package core

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) UpdateTx(ctx context.Context, oldTxId string, newTxId string, filter model.FileFilter) error {
	tx := u.txStore.Delete(oldTxId)
	if tx == nil {
		return nil
	}

	tx.Lock()
	defer func() {
		tx.Unlock()
		// TODO free tx
	}()

	newTx, ok := u.txStore.Get(newTxId)
	if !ok {
		newTx = &core.Transaction{} // TODO use pool
		u.txStore.Put(newTxId, newTx)
	}

	newTx.RLock()
	var (
		err error

		files       = make([]model.File, 0, tx.Len())
		freeNodes   = make([]*core.Node[model.File], 0, tx.Len())
		deleteFiles = make([]model.File, 0, tx.Len())
	)
	defer func() {
		u.m.Lock()
		u.deleteFiles = append(u.deleteFiles, deleteFiles...)
		u.m.Unlock()
	}()

	for key, f := range tx.Files() {
		if filter.BeforeSeq != nil && newTx.File(key).Latest().Seq.After(*filter.BeforeSeq) {
			err = fs_db.TxSerializationErr
		}

		n := f.PopBack()
		if n == nil {
			continue
		}

		file := n.V()
		file.TxId = newTxId
		file.Seq = sequence.Next()
		files = append(files, file)
		freeNodes = append(freeNodes, n)

		for n = f.PopFront(); n != nil; n = f.PopFront() {
			deleteFiles = append(deleteFiles, n.V())
			freeNodes = append(freeNodes, n)
		}
	}
	newTx.RUnlock()
	if err != nil {
		deleteFiles = append(deleteFiles, files...)
		return err
	}

	if len(files) == 0 {
		return nil
	}

	err = u.fileRepo.RunTransaction(ctx, func(ctx context.Context) error {
		for _, f := range files {
			err = u.fileRepo.Set(ctx, f)
			if err != nil {
				return fmt.Errorf("store to tx: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("run transaction: %w", err)
	}

	newTx.Lock()
	u.allStore.Lock()
	for _, f := range files {
		u.storeToTx(newTx, f)
	}
	newTx.Unlock()

	for _, n := range freeNodes {
		n.DeleteLink() // TODO free link node
		// TODO free node
	}
	u.allStore.Unlock()

	return nil
}
