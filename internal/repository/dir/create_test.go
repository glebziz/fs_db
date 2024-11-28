package dir

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Create(t *testing.T) {
	var (
		rootPath = testNewRootPath(t)

		dir     = gofakeit.UUID()
		dirPath = path.Join(rootPath, dir)
	)

	r, err := New([]string{rootPath})
	require.NoError(t, err)
	require.NotNil(t, r)

	err = r.Create(context.Background(), model.Dir{
		Name: dir,
		Root: rootPath,
	})
	require.NoError(t, err)
	require.EqualValues(t, 1, r.counts[rootPath])
	require.Equal(t, model.Dir{
		Name: dir,
		Root: rootPath,
	}, r.dirs[dirPath])

	st, err := os.Stat(dirPath)
	require.NoError(t, err)
	require.Equal(t, true, st.IsDir())
}
