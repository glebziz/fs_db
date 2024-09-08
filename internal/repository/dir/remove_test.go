package dir

import (
	"context"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Remove(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var (
			rootPath = testNewRootPath(t)
			dir      = gofakeit.UUID()
		)

		testCreateDir(t, path.Join(rootPath, dir))

		r, err := New([]string{rootPath})
		require.NoError(t, err)
		require.NotNil(t, r)
		require.EqualValues(t, 1, r.counts[rootPath])

		err = r.Remove(context.Background(), model.Dir{
			Name: dir,
			Root: rootPath,
		})
		require.NoError(t, err)
		require.Zero(t, r.counts[rootPath])
	})

	t.Run("success without dir", func(t *testing.T) {
		var (
			rootPath = testNewRootPath(t)
			dir      = gofakeit.UUID()
		)

		testCreateDir(t, path.Join(rootPath, dir))

		r, err := New([]string{rootPath})
		require.NoError(t, err)
		require.NotNil(t, r)
		require.EqualValues(t, 1, r.counts[rootPath])

		err = r.Remove(context.Background(), model.Dir{
			Name: gofakeit.UUID(),
			Root: rootPath,
		})
		require.NoError(t, err)
		require.EqualValues(t, 1, r.counts[rootPath])
	})
}
