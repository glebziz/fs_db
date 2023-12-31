package dir

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_CountByParent(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r := New(manager)

		err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
			var (
				parent1 = gofakeit.UUID()
				parent2 = gofakeit.UUID()
				dirs    = []model.Dir{{
					Id:         gofakeit.UUID(),
					ParentPath: parent1,
				}, {
					Id:         gofakeit.UUID(),
					ParentPath: parent2,
				}, {
					Id:         gofakeit.UUID(),
					ParentPath: parent1,
				}, {
					Id:         gofakeit.UUID(),
					ParentPath: parent1,
				}}
			)

			for _, dir := range dirs {
				testCreateDir(ctx, t, &dir)
			}

			actual, err := r.CountByParent(ctx, parent1)
			require.NoError(t, err)
			require.Equal(t, uint64(3), actual)

			actual, err = r.CountByParent(ctx, parent2)
			require.NoError(t, err)
			require.Equal(t, uint64(1), actual)

			return assert.AnError
		})

		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("empty dirs", func(t *testing.T) {
		t.Parallel()

		r := New(manager)

		err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
			actual, err := r.CountByParent(ctx, gofakeit.UUID())
			require.NoError(t, err)
			require.Zero(t, actual)

			return assert.AnError
		})

		require.ErrorIs(t, err, assert.AnError)
	})
}
