package transaction

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) Begin(ctx context.Context, isoLevel model.TxIsoLevel) (string, error) {
	id := u.idGen.Generate()

	err := u.txRepo.Store(ctx, model.Transaction{
		Id:       id,
		IsoLevel: isoLevel,
		Seq:      sequence.Next(),
	})
	if err != nil {
		return "", fmt.Errorf("tx repository store: %w", err)
	}

	return id, nil
}
