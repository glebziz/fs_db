package dir

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func testNewRootPath(t *testing.T) string {
	t.Helper()

	rootPath, err := os.MkdirTemp("", "dir_rep")
	require.NoError(t, err)
	t.Cleanup(func() {
		err = os.RemoveAll(rootPath)
		require.NoError(t, err)
	})

	return rootPath
}

func testCreateDir(t *testing.T, path string) {
	t.Helper()

	err := os.Mkdir(path, mkdirPerm)
	require.NoError(t, err)
}

func testCreateFile(t *testing.T, path string) {
	t.Helper()

	f, err := os.Create(path)
	require.NoError(t, err)

	err = f.Close()
	require.NoError(t, err)
}
