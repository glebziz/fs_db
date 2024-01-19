package file

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Delete_Success(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		txId = gofakeit.UUID()
		key1 = gofakeit.UUID()
		key2 = gofakeit.UUID()
	)

	err := r.Delete(ctx, txId, key1)
	require.NoError(t, err)

	err = r.Delete(ctx, txId, key2)
	require.NoError(t, err)

	actual := testGetFile(ctx, t, r.p, key1)
	require.Equal(t, &model.File{
		Key: key1,
	}, actual)

	actual = testGetFile(ctx, t, r.p, key2)
	require.Equal(t, &model.File{
		Key: key2,
	}, actual)

}

func TestRep_Delete_Error(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	err := r.Delete(ctx, gofakeit.UUID(), "")
	require.Error(t, err)
}
