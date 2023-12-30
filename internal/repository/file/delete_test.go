package dir

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func TestRep_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r := New(manager)

		err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
			var (
				file = model.File{
					Id:         gofakeit.UUID(),
					Key:        gofakeit.UUID(),
					ParentPath: gofakeit.UUID(),
				}
			)

			testCreateFile(ctx, t, &file)

			err := r.Delete(ctx, file.Key)
			require.NoError(t, err)

			return assert.AnError
		})

		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("success with non existing file", func(t *testing.T) {
		t.Parallel()

		r := New(manager)

		err := r.RunTransaction(context.Background(), func(ctx context.Context) error {
			err := r.Delete(ctx, gofakeit.UUID())
			require.ErrorIs(t, err, pkgModel.NotFoundErr)

			return assert.AnError
		})

		require.ErrorIs(t, err, assert.AnError)
	})
}
