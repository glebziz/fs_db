package transaction

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Oldest_Success(t *testing.T) {
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

	actual, err := r.Oldest(testCtx)
	require.NoError(t, err)
	require.Equal(t, &txs[0], actual)
}

func TestRep_Oldest_Error(t *testing.T) {
	t.Parallel()

	r := New()

	actual, err := r.Oldest(testCtx)
	require.ErrorIs(t, err, fs_db.TxNotFoundErr)
	require.Nil(t, actual)
}
