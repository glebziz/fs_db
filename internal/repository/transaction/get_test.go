package transaction

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Get_Success(t *testing.T) {
	t.Parallel()

	r := New()

	var (
		txs = []model.Transaction{{
			Id:  gofakeit.UUID(),
			Seq: 1,
		}, {
			Id:  gofakeit.UUID(),
			Seq: 2,
		}}
	)

	for _, tx := range txs {
		testCreateTransaction(t, r, tx)
	}

	actual, err := r.Get(context.Background(), model.MainTxId)
	require.NoError(t, err)
	require.Equal(t, model.Transaction{
		Id:       model.MainTxId,
		IsoLevel: fs_db.IsoLevelDefault,
	}, actual)

	for _, tx := range txs {
		actual, err = r.Get(context.Background(), tx.Id)
		require.NoError(t, err)
		require.Equal(t, tx, actual)
	}
}

func TestRep_Get_Error(t *testing.T) {
	t.Parallel()

	r := New()

	actual, err := r.Get(context.Background(), gofakeit.UUID())
	require.ErrorIs(t, err, fs_db.ErrTxNotFound)
	require.Equal(t, model.Transaction{}, actual)
}
