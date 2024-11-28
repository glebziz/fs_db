package dir

import (
	"context"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/disk"
)

func TestRep_Get(t *testing.T) {
	var (
		rootPath = testNewRootPath(t)

		dir1 = gofakeit.UUID()
		dir2 = gofakeit.UUID()

		file1 = path.Join(rootPath, dir1, gofakeit.UUID())
		file2 = path.Join(rootPath, dir1, gofakeit.UUID())
		file3 = path.Join(rootPath, dir2, gofakeit.UUID())
	)

	testCreateDir(t, path.Join(rootPath, dir1))
	testCreateDir(t, path.Join(rootPath, dir2))
	testCreateFile(t, file1)
	testCreateFile(t, file2)
	testCreateFile(t, file3)

	r, err := New([]string{rootPath})
	require.NoError(t, err)
	require.NotNil(t, r)

	st, err := disk.Usage(context.Background(), rootPath)
	require.NoError(t, err)
	require.NotNil(t, st)

	dirs, err := r.Get(context.Background())
	require.NoError(t, err)
	require.Len(t, dirs, 2)

	require.True(t, gomock.InAnyOrder(model.Dirs{{
		Name:  dir1,
		Root:  rootPath,
		Count: 2,
		Free:  st.Free,
	}, {
		Name:  dir2,
		Root:  rootPath,
		Count: 1,
		Free:  st.Free,
	}}).Matches(dirs))
}
