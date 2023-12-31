package dir

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Store_Success(t *testing.T) {
	t.Parallel()

	r := New(manager)

	err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
		var (
			file1 = model.File{
				Id:         gofakeit.UUID(),
				Key:        gofakeit.UUID(),
				ParentPath: gofakeit.UUID(),
			}

			file2 = model.File{
				Id:         gofakeit.UUID(),
				Key:        gofakeit.UUID(),
				ParentPath: gofakeit.UUID(),
			}
			file3 = model.File{
				Id:         gofakeit.UUID(),
				Key:        file1.Key,
				ParentPath: gofakeit.UUID(),
			}
		)

		err := r.Store(ctx, file1)
		require.NoError(t, err)

		err = r.Store(ctx, file2)
		require.NoError(t, err)

		actual := testGetFile(ctx, t, file1.Key)
		require.Equal(t, &file1, actual)

		actual = testGetFile(ctx, t, file2.Key)
		require.Equal(t, &file2, actual)

		err = r.Store(ctx, file3)
		require.NoError(t, err)

		actual = testGetFile(ctx, t, file1.Key)
		require.Equal(t, &file3, actual)

		return assert.AnError
	})

	require.ErrorIs(t, err, assert.AnError)
}

func TestRep_Create_Error(t *testing.T) {
	t.Parallel()

	r := New(manager)

	err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
		err := r.Store(ctx, model.File{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		})
		require.Error(t, err)

		err = r.Store(ctx, model.File{
			Id:  gofakeit.UUID(),
			Key: gofakeit.UUID(),
		})
		require.Error(t, err)

		return assert.AnError
	})

	require.ErrorIs(t, err, assert.AnError)
}
