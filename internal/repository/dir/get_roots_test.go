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

func TestRepo_GetRoots(t *testing.T) {
	var (
		rootPath1 = testNewRootPath(t)
		rootPath2 = testNewRootPath(t)
		rootPath3 = testNewRootPath(t)

		dir1 = path.Join(rootPath1, gofakeit.UUID())
		dir2 = path.Join(rootPath1, gofakeit.UUID())
		dir3 = path.Join(rootPath2, gofakeit.UUID())
	)

	testCreateDir(t, dir1)
	testCreateDir(t, dir2)
	testCreateDir(t, dir3)

	r, err := New(context.Background(), []string{rootPath1, rootPath2, rootPath3}, osa.Adapter{})
	require.NoError(t, err)
	require.NotNil(t, r)

	roots, err := r.GetRoots(context.Background())
	require.NoError(t, err)
	require.Len(t, roots, 3)
	require.Equal(t, []model.Root{{
		Path:  rootPath1,
		Count: 2,
	}, {
		Path:  rootPath2,
		Count: 1,
	}, {
		Path: rootPath3,
	}}, roots)
}
