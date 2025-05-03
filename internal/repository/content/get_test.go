package content

import (
	"context"
	"io"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	osa "github.com/glebziz/fs_db/internal/adapter/os"
)

func TestRepo_Get(t *testing.T) {
	t.Parallel()

	const (
		filePath = "filePath"
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					Open(gomock.Any(), filePath).
					Return(nil, nil)
			},
		},
		{
			name: "ErrNotExist error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					Open(gomock.Any(), gomock.Any()).
					Return(nil, os.ErrNotExist)
			},
			err: fs_db.ErrNotFound,
		},
		{
			name: "Open error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					Open(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			r := td.newRepo()
			f, err := r.Get(context.Background(), filePath)

			require.ErrorIs(t, err, tc.err)
			require.Nil(t, f)
		})
	}
}

func TestRepo_Get_Success_Int(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := New(osa.Adapter{})

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0750)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())
		testCreateFile(t, fPath, content)

		c, err := r.Get(context.Background(), fPath)

		require.NoError(t, err)

		actContent, err := io.ReadAll(c)
		require.NoError(t, err)
		require.Equal(t, content, actContent)

		err = c.Close()
		require.NoError(t, err)
	})

	t.Run("get error", func(t *testing.T) {
		r := New(osa.Adapter{})

		c, err := r.Get(context.Background(), path.Join(rootPath, gofakeit.UUID()))

		require.ErrorIs(t, err, fs_db.ErrNotFound)
		require.Nil(t, c)
	})
}
