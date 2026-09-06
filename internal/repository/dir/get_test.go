package dir

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	osa "github.com/glebziz/fs_db/internal/adapter/os"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRepo_Get(t *testing.T) {
	t.Parallel()

	const (
		free uint64 = 100

		dirName1 = "dirName1"
		dirName2 = "dirName2"
		rootDir  = "rootDir"
	)

	for _, tc := range []struct {
		name    string
		prepare func(td *testDeps, r *Repo)
		dirs    model.Dirs
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps, r *Repo) {
				r.counts[rootDir] = 2
				r.dirs[path.Join(rootDir, dirName1)] = model.Dir{
					Name: dirName1,
					Root: rootDir,
				}
				r.dirs[path.Join(rootDir, dirName2)] = model.Dir{
					Name: dirName2,
					Root: rootDir,
				}

				td.os.EXPECT().
					Usage(gomock.Any(), rootDir).
					Return(&disk.UsageStat{
						Free: free,
					}, nil)

				td.os.EXPECT().
					ReadDir(context.Background(), path.Join(rootDir, dirName1)).
					Return([]os.DirEntry{dirEntry{}, dirEntry{}}, nil)

				td.os.EXPECT().
					ReadDir(context.Background(), path.Join(rootDir, dirName2)).
					Return([]os.DirEntry{dirEntry{}}, nil)
			},
			dirs: model.Dirs{{
				Name:  dirName1,
				Root:  rootDir,
				Free:  free,
				Count: 2,
			}, {
				Name:  dirName2,
				Root:  rootDir,
				Free:  free,
				Count: 1,
			}},
		},
		{
			name: "Usage error",
			prepare: func(td *testDeps, r *Repo) {
				r.counts[rootDir] = 2
				r.dirs[path.Join(rootDir, dirName1)] = model.Dir{
					Name: dirName1,
					Root: rootDir,
				}
				r.dirs[path.Join(rootDir, dirName2)] = model.Dir{
					Name: dirName2,
					Root: rootDir,
				}

				td.os.EXPECT().
					Usage(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "ReadDir error",
			prepare: func(td *testDeps, r *Repo) {
				r.counts[rootDir] = 2
				r.dirs[path.Join(rootDir, dirName1)] = model.Dir{
					Name: dirName1,
					Root: rootDir,
				}
				r.dirs[path.Join(rootDir, dirName2)] = model.Dir{
					Name: dirName2,
					Root: rootDir,
				}

				td.os.EXPECT().
					Usage(gomock.Any(), gomock.Any()).
					Return(&disk.UsageStat{
						Free: free,
					}, nil)

				td.os.EXPECT().
					ReadDir(context.Background(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t, []string{rootDir})
			r := td.newRepo()

			tc.prepare(td, r)

			dirs, err := r.Get(context.Background())

			require.ErrorIs(t, err, tc.err)
			require.True(t, gomock.InAnyOrder(dirs).Matches(tc.dirs))
		})
	}
}

func TestRepo_Get_Int(t *testing.T) {
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

	r, err := New(context.Background(), []string{rootPath}, osa.Adapter{})
	require.NoError(t, err)
	require.NotNil(t, r)

	st, err := disk.UsageWithContext(context.Background(), rootPath)
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
