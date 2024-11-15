package core

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) Load(ctx context.Context) ([]model.File, error) {
	files, err := u.fileRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("db get all: %w", err)
	}

	var (
		deleteFiles = make([]model.File, 0, len(files))
		mainFiles   = make(map[string]model.File, len(files))
	)
	for _, file := range files {
		if file.TxId != model.MainTxId {
			deleteFiles = append(deleteFiles, file)
			continue
		}

		mainFile, ok := mainFiles[file.Key]
		if file.Seq.Before(mainFile.Seq) {
			deleteFiles = append(deleteFiles, file)
			continue
		}

		mainFiles[file.Key] = file
		if ok {
			deleteFiles = append(deleteFiles, mainFile)
		}
	}

	var (
		maxSeq = sequence.Seq(1)
		mainTx = &core.Transaction{}
	)
	u.txStore.Put(model.MainTxId, mainTx)

	for _, file := range mainFiles {
		maxSeq = max(maxSeq, file.Seq)
		u.storeToTx(mainTx, file)
	}

	sequence.Set(maxSeq)

	return deleteFiles, nil
}
