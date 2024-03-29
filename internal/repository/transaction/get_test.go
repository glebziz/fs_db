package transaction

import (
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
			Id:       gofakeit.UUID(),
			CreateTs: gofakeit.Date(),
		}, {
			Id:       gofakeit.UUID(),
			CreateTs: gofakeit.Date(),
		}}
	)

	for _, tx := range txs {
		testCreateTransaction(t, r, tx)
	}

	actual, err := r.Get(testCtx, model.MainTxId)
	require.NoError(t, err)
	require.Equal(t, &model.Transaction{
		Id:       model.MainTxId,
		IsoLevel: fs_db.IsoLevelDefault,
	}, actual)

	for _, tx := range txs {
		actual, err = r.Get(testCtx, tx.Id)
		require.NoError(t, err)
		require.Equal(t, &tx, actual)
	}
}

func TestRep_Get_Error(t *testing.T) {
	t.Parallel()

	r := New()

	actual, err := r.Get(testCtx, gofakeit.UUID())
	require.ErrorIs(t, err, fs_db.TxNotFoundErr)
	require.Nil(t, actual)
}
