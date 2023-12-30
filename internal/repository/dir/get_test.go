package dir

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r := New(manager)

		err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
			var (
				dirs = []model.Dir{{
					Id:         gofakeit.UUID(),
					FileCount:  0,
					ParentPath: gofakeit.UUID(),
				}, {
					Id:         gofakeit.UUID(),
					FileCount:  3,
					ParentPath: gofakeit.UUID(),
				}}

				files = []model.File{{
					Id:         gofakeit.UUID(),
					Key:        gofakeit.UUID(),
					ParentPath: dirs[1].GetPath(),
				}, {
					Id:         gofakeit.UUID(),
					Key:        gofakeit.UUID(),
					ParentPath: dirs[1].GetPath(),
				}, {
					Id:         gofakeit.UUID(),
					Key:        gofakeit.UUID(),
					ParentPath: dirs[1].GetPath(),
				}}
			)

			for _, file := range files {
				testCreateFile(ctx, t, &file)
			}

			for _, dir := range dirs {
				testCreateDir(ctx, t, &dir)
			}

			actual, err := r.Get(ctx)
			require.NoError(t, err)
			require.Equal(t, dirs, actual)

			return assert.AnError
		})

		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("empty dirs", func(t *testing.T) {
		t.Parallel()

		r := New(manager)

		err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
			actual, err := r.Get(ctx)
			require.NoError(t, err)
			require.Empty(t, actual)

			return assert.AnError
		})

		require.ErrorIs(t, err, assert.AnError)
	})
}
