package core

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) UpdateTx(ctx context.Context, oldTxId, newTxId string, filter model.FileFilter) (deleteFiles []model.File, err error) { //nolint:funlen,cyclop,lll // TODO fix
	tx := u.txStore.Delete(oldTxId)
	if tx == nil {
		return nil, nil
	}

	tx.Lock()
	defer func() {
		tx.Unlock()
		u.txPool.Release(tx)
	}()

	newTx, ok := u.txStore.Get(newTxId)
	if !ok {
		newTx = u.txPool.Acquire()
		u.txStore.Put(newTxId, newTx)
	}

	newTx.RLock()
	var (
		files     = make([]model.File, 0, tx.Len())
		freeNodes = make([]*core.Node[model.File], 0, tx.Len())
	)
	defer func() {
		for _, n := range freeNodes {
			link := n.DeleteLink()
			u.nodePool.Release(link, n)
		}
		deleteFiles = append(deleteFiles, files...)
	}()

	deleteFiles = make([]model.File, 0, tx.Len())
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
		return
	}

	if len(files) == 0 {
		return
	}

	newTx.Lock()
	u.allStore.Lock()
	defer func() {
		u.allStore.Unlock()
		newTx.Unlock()
	}()

	err = u.fileRepo.RunTransaction(ctx, func(ctx context.Context) error {
		for i := range files {
			files[i].Seq = sequence.Next()
			err = u.fileRepo.Set(ctx, files[i])
			if err != nil {
				return fmt.Errorf("store to tx: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return
	}

	for _, f := range files {
		u.storeToTx(newTx, f)
	}

	files = nil
	return
}
