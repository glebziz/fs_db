package file

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Store_Success(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		txId  = gofakeit.UUID()
		file1 = model.File{
			Key:       gofakeit.UUID(),
			ContentId: gofakeit.UUID(),
		}
		file2 = model.File{
			Key:       gofakeit.UUID(),
			ContentId: gofakeit.UUID(),
		}
	)

	err := r.Store(ctx, txId, file1)
	require.NoError(t, err)

	err = r.Store(ctx, txId, file2)
	require.NoError(t, err)

	actual := testGetFile(ctx, t, r.p, file1.Key)
	require.Equal(t, &file1, actual)

	actual = testGetFile(ctx, t, r.p, file2.Key)
	require.Equal(t, &file2, actual)
}

func TestRep_Create_Error(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	err := r.Store(ctx, gofakeit.UUID(), model.File{
		ContentId: gofakeit.UUID(),
	})
	require.Error(t, err)
}
