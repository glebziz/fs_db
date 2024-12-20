package transaction

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Delete_Success(t *testing.T) {
	t.Parallel()

	r := New()

	var (
		tx = model.Transaction{
			Id:  gofakeit.UUID(),
			Seq: 1,
		}
	)

	testCreateTransaction(t, r, tx)

	actual, err := r.Delete(context.Background(), tx.Id)
	require.NoError(t, err)
	require.Equal(t, tx, actual)
}

func TestRep_Delete_Error(t *testing.T) {
	t.Parallel()

	r := New()

	actual, err := r.Delete(context.Background(), gofakeit.UUID())
	require.ErrorIs(t, err, fs_db.ErrTxNotFound)
	require.Equal(t, model.Transaction{}, actual)
}
