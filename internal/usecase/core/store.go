package core

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) Store(ctx context.Context, f model.File) error {

	tx, ok := u.txStore.Get(f.TxId)
	if !ok {
		tx = &core.Transaction{}
		u.txStore.Put(f.TxId, tx)
	}

	tx.Lock()
	u.allStore.Lock()
	defer func() {
		u.allStore.Unlock()
		tx.Unlock()
	}()

	f.Seq = sequence.Next()
	err := u.fileRepo.Set(ctx, f)
	if err != nil {
		return fmt.Errorf("db set: %w", err)
	}

	u.storeToTx(tx, f)

	return nil
}

func (u *useCase) storeToTx(tx *core.Transaction, f model.File) {
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
