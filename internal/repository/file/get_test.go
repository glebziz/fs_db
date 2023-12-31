package dir

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Get_Success(t *testing.T) {
	t.Parallel()

	r := New(manager)

	err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
		var (
			files = []model.File{{
				Id:         gofakeit.UUID(),
				Key:        gofakeit.UUID(),
				ParentPath: gofakeit.UUID(),
			}, {
				Id:         gofakeit.UUID(),
				Key:        gofakeit.UUID(),
				ParentPath: gofakeit.UUID(),
			}}
		)

		for _, file := range files {
			testCreateFile(ctx, t, &file)
		}

		for _, file := range files {
			actual, err := r.Get(ctx, file.Key)
			require.NoError(t, err)
			require.Equal(t, &file, actual)
		}

		return assert.AnError
	})

	require.ErrorIs(t, err, assert.AnError)
}

func TestRep_Get_Error(t *testing.T) {
	t.Parallel()

	r := New(manager)

	err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
		actual, err := r.Get(ctx, gofakeit.UUID())
		require.ErrorIs(t, err, fs_db.NotFoundErr)
		require.Nil(t, actual)

		return assert.AnError
	})

	require.ErrorIs(t, err, assert.AnError)
}
