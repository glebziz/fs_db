package dir

import (
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestNew(t *testing.T) {
	var (
		rootPath  = testNewRootPath(t)
		rootPath2 = testNewRootPath(t)

		dirName1 = gofakeit.UUID()
		dirName2 = "1234567890"
		dirName3 = gofakeit.UUID()

		dir1 = path.Join(rootPath, dirName1)
		dir2 = path.Join(rootPath, dirName2)
		dir3 = path.Join(rootPath, dirName3)
		file = path.Join(rootPath, gofakeit.UUID())
	)

	testCreateDir(t, dir1)
	testCreateDir(t, dir2)
	testCreateDir(t, dir3)
	testCreateFile(t, file)

	err := os.Remove(rootPath2)
	require.NoError(t, err)

	r, err := New([]string{rootPath, rootPath2})

	require.NoError(t, err)
	require.Equal(t, []string{rootPath, rootPath2}, r.roots)
	require.EqualValues(t, 2, r.counts[rootPath])
	require.EqualValues(t, 0, r.counts[rootPath2])
	require.Equal(t, map[string]model.Dir{
		dir1: {
			Name: dirName1,
			Root: rootPath,
		},
		dir3: {
			Name: dirName3,
			Root: rootPath,
		},
	}, r.dirs)
}
