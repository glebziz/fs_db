package core

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) GetFiles(_ context.Context, txId string, filter model.FileFilter) ([]model.File, error) {
	var (
		files  []model.File
		sFiles []model.File
	)
	if filter.BeforeSeq == nil && filter.TxId == nil {
		files = u.getFilesFromTx(&u.allStore, nil)
	} else if filter.TxId != nil {
		tx, ok := u.txStore.Get(txId)
		if ok {
			files = u.getFilesFromTx(tx, nil)
		}

		sTx, ok := u.txStore.Get(*filter.TxId)
		if ok {
			sFiles = u.getFilesFromTx(sTx, filter.BeforeSeq)
		}
	}

	return u.mergeFiles(files, sFiles), nil
}

func (*useCase) getFilesFromTx(tx *core.Transaction, beforeSeq *sequence.Seq) []model.File {
	tx.RLock()
	defer tx.RUnlock()

	files := make([]model.File, 0, tx.Len())
	for _, f := range tx.Files() {
		if beforeSeq == nil {
			files = append(files, f.Latest())
		} else {
			files = append(files, f.LastBefore(*beforeSeq))
		}
	}

	return files
}

func (*useCase) mergeFiles(files, sFiles []model.File) []model.File {
	if len(sFiles) > len(files) {
		sFiles, files = files, sFiles
	}

	filesByKey := make(map[string]model.File, len(files))
	for _, file := range files {
		filesByKey[file.Key] = file
	}
	for _, file := range sFiles {
		filesByKey[file.Key] = file.Latest(filesByKey[file.Key])
	}

	files = files[:0]
	for _, file := range filesByKey {
		if file.Seq.Zero() {
			continue
		}

		files = append(files, file)
	}

	return files
}
