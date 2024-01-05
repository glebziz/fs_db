package transaction

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Store_Success(t *testing.T) {
	t.Parallel()

	r := New()

	var (
		tx1 = model.Transaction{
			Id:       gofakeit.UUID(),
			CreateTs: gofakeit.Date(),
		}
		tx2 = model.Transaction{
			Id:       gofakeit.UUID(),
			CreateTs: gofakeit.Date(),
		}
	)

	err := r.Store(testCtx, tx1)
	require.NoError(t, err)

	err = r.Store(testCtx, tx2)
	require.NoError(t, err)

	actual := testGetTransaction(t, r, tx1.Id)
	require.Equal(t, &tx1, actual)

	actual = testGetTransaction(t, r, tx2.Id)
	require.Equal(t, &tx2, actual)
}

func TestRep_Create_Error(t *testing.T) {
	t.Parallel()

	r := New()

	var (
		tx = model.Transaction{
			Id:       gofakeit.UUID(),
			CreateTs: gofakeit.Date(),
		}
	)

	err := r.Store(testCtx, tx)
	require.NoError(t, err)

	err = r.Store(testCtx, model.Transaction{
		Id:       tx.Id,
		CreateTs: gofakeit.Date(),
	})
	require.ErrorIs(t, err, fs_db.TxAlreadyExistsErr)
}
