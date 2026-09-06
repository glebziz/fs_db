package dir

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/repository/dir/mocks"
)

type testDeps struct {
	t *testing.T

	rootDirs []string

	os *mocks.MockosAdapter
}

func newTestDeps(t *testing.T, rootDirs []string) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		t:        t,
		rootDirs: rootDirs,
		os:       mocks.NewMockosAdapter(ctrl),
	}
}

func (td *testDeps) newRepo() *Repo {
	return &Repo{
		roots:  td.rootDirs,
		dirs:   make(map[string]model.Dir),
		counts: make(map[string]uint64),
		os:     td.os,
	}
}

type dirEntry struct {
	name string
}

func (de dirEntry) Name() string {
	return de.name
}

func (de dirEntry) IsDir() bool {
	return true
}

func (de dirEntry) Type() fs.FileMode {
	return fs.ModeDir
}

func (de dirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

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
