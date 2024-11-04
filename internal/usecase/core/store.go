package core

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) Store(ctx context.Context, f model.File) error {
	err := u.fileRepo.Set(ctx, f)
	if err != nil {
		return fmt.Errorf("db set: %w", err)
	}

	tx, ok := u.txStore.Get(f.TxId)
	if !ok {
		tx = &core.Transaction{}
		u.txStore.Put(f.TxId, tx)
	}

	tx.Lock()
	u.allStore.Lock()
	u.storeToTx(tx, f)
	u.allStore.Unlock()
	tx.Unlock()

	return nil
}

func (u *useCase) storeToTx(tx *core.Transaction, f model.File) {
	f.Seq = sequence.Next()

	var (
		link = (&core.Node[model.File]{}).
			SetV(f)
		n = (&core.Node[model.File]{}).
			SetV(f).
			SetLink(link)
	)

	tx.PushBack(n)
	u.allStore.PushBack(link)
}
