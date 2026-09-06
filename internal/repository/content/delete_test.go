package content

import (
	"context"
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

func TestRepo_Delete(t *testing.T) {
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
					Remove(gomock.Any(), filePath).
					Return(nil)
			},
		},
		{
			name: "ErrNotExist error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					Remove(gomock.Any(), gomock.Any()).
					Return(os.ErrNotExist)
			},
			err: fs_db.ErrNotFound,
		},
		{
			name: "Remove error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					Remove(gomock.Any(), gomock.Any()).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			r := td.newRepo()
			err := r.Delete(context.Background(), filePath)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRepo_Delete_Int(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := New(osa.Adapter{})

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0750)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())
		testCreateFile(t, fPath, content)

		err = r.Delete(context.Background(), fPath)

		require.NoError(t, err)
	})

	t.Run("success with non existing file", func(t *testing.T) {
		r := New(osa.Adapter{})

		dir := path.Join(rootPath, gofakeit.UUID())
		fPath := path.Join(dir, gofakeit.UUID())

		err := r.Delete(context.Background(), fPath)
		require.ErrorIs(t, err, fs_db.ErrNotFound)
	})
}
