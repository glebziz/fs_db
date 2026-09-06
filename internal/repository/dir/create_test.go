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

func TestRepo_Create(t *testing.T) {
	t.Parallel()

	const (
		dirName = "dirName"
		rootDir = "rootDir"
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
					MkdirAll(gomock.Any(), path.Join(rootDir, dirName), mkdirPerm).
					Return(nil)
			},
			checkR: func(t *testing.T, r *Repo) {
				require.Equal(t, model.Dir{
					Name: dirName,
					Root: rootDir,
				}, r.dirs[path.Join(rootDir, dirName)])
				require.Equal(t, uint64(1), r.counts[rootDir])
			},
		},
		{
			name: "MkdirAll error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					MkdirAll(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)
			},
			err:    assert.AnError,
			checkR: func(t *testing.T, r *Repo) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t, []string{rootDir})
			tc.prepare(td)

			r := td.newRepo()
			err := r.Create(context.Background(), model.Dir{
				Name: dirName,
				Root: rootDir,
			})
			require.ErrorIs(t, err, tc.err)
			tc.checkR(t, r)
		})
	}
}

func TestRepo_Create_Int(t *testing.T) {
	var (
		rootPath = testNewRootPath(t)

		dir     = gofakeit.UUID()
		dirPath = path.Join(rootPath, dir)
	)

	r, err := New(context.Background(), []string{rootPath}, osa.Adapter{})
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
