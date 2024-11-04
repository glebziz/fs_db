//go:build inline

package db

import (
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/config"
	inlineDb "github.com/glebziz/fs_db/pkg/inline/db"
)

func newTestDb(t testing.TB) fs_db.DB {
	t.Helper()

	dir, err := os.MkdirTemp("", "fs_db_test")
	require.NoError(t, err)

	_db, err := inlineDb.New(testCtx, &config.Storage{
		MaxDirCount: 100,
		DbPath:      path.Join(dir, "test_db"),
		RootDirs:    []string{path.Join(dir, gofakeit.UUID()), path.Join(dir, gofakeit.UUID())},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = _db.Close()
		require.NoError(t, err)

		err = os.RemoveAll(dir)
		require.NoError(t, err)
	})

	return _db
}
