package dir

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	osa "github.com/glebziz/fs_db/internal/adapter/os"
	"github.com/glebziz/fs_db/internal/model"
)

func TestNew(t *testing.T) {
	t.Parallel()

	const (
		dirName1 = "00000000-0000-0000-0000-000000000000"
		dirName2 = "00000000-0000-0000-0000-000000000001"
		dirName3 = "dirName"
		rootDir1 = "rootDir1"
		rootDir2 = "rootDir2"
	)

	for _, tc := range []struct {
		name    string
		prepare func(td *testDeps)
		err     error
		checkR  func(t *testing.T, r *Repo)
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					ReadDir(gomock.Any(), rootDir1).
					Return([]os.DirEntry{dirEntry{
						name: dirName1,
					}, dirEntry{
						name: dirName2,
					}, dirEntry{
						name: dirName3,
					}}, nil)

				td.os.EXPECT().
					ReadDir(gomock.Any(), rootDir2).
					Return(nil, os.ErrNotExist)

				td.os.EXPECT().
					MkdirAll(gomock.Any(), rootDir2, mkdirPerm).
					Return(nil)
			},
			checkR: func(t *testing.T, r *Repo) {
				require.Equal(t, []string{rootDir1, rootDir2}, r.roots)
				require.Equal(t, uint64(2), r.counts[rootDir1])
				require.Zero(t, r.counts[rootDir2])
				require.Equal(t, model.Dir{
					Name: dirName1,
					Root: rootDir1,
				}, r.dirs[path.Join(rootDir1, dirName1)])
				require.Equal(t, model.Dir{
					Name: dirName2,
					Root: rootDir1,
				}, r.dirs[path.Join(rootDir1, dirName2)])
			},
		},
		{
			name: "ReadDir error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					ReadDir(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
			checkR: func(t *testing.T, r *Repo) {
				require.Nil(t, r)
			},
		},
		{
			name: "MkdirAll error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					ReadDir(gomock.Any(), gomock.Any()).
					Return(nil, os.ErrNotExist)

				td.os.EXPECT().
					MkdirAll(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)
			},
			err: assert.AnError,
			checkR: func(t *testing.T, r *Repo) {
				require.Nil(t, r)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t, nil)
			tc.prepare(td)

			r, err := New(context.Background(), []string{"./" + rootDir1, rootDir2}, td.os)

			require.ErrorIs(t, err, tc.err)
			tc.checkR(t, r)
		})
	}
}

func TestNew_Int(t *testing.T) {
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

	r, err := New(context.Background(), []string{rootPath, rootPath2}, osa.Adapter{})

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
