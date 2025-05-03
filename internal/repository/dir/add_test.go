package dir

import (
	"context"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	osa "github.com/glebziz/fs_db/internal/adapter/os"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRepo_Add(t *testing.T) {
	t.Run("success without dir", func(t *testing.T) {
		var (
			rootPath = testNewRootPath(t)
			dir      = gofakeit.UUID()
		)

		testCreateDir(t, path.Join(rootPath, dir))

		r, err := New(context.Background(), []string{rootPath}, osa.Adapter{})
		require.NoError(t, err)
		require.NotNil(t, r)
		require.EqualValues(t, 1, r.counts[rootPath])

		err = r.Add(context.Background(), model.Dir{
			Name: gofakeit.UUID(),
			Root: rootPath,
		})
		require.NoError(t, err)
		require.EqualValues(t, 2, r.counts[rootPath])
	})

	t.Run("success with dir already exists", func(t *testing.T) {
		var (
			rootPath = testNewRootPath(t)
			dir      = gofakeit.UUID()
		)

		testCreateDir(t, path.Join(rootPath, dir))

		r, err := New(context.Background(), []string{rootPath}, osa.Adapter{})
		require.NoError(t, err)
		require.NotNil(t, r)
		require.EqualValues(t, 1, r.counts[rootPath])

		err = r.Add(context.Background(), model.Dir{
			Name: dir,
			Root: rootPath,
		})
		require.NoError(t, err)
		require.EqualValues(t, 1, r.counts[rootPath])
	})
}
