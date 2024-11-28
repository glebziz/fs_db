package content

import (
	"bytes"
	"context"
	"io"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/os"
)

func TestRep_Store(t *testing.T) {
	t.Run("success with existing dir", func(t *testing.T) {
		r := New()

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0750)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())

		err = r.Store(context.Background(), fPath, bytes.NewReader(content))
		require.NoError(t, err)
	})

	t.Run("success with existing file", func(t *testing.T) {
		r := New()

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0750)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())
		testCreateFile(t, fPath, content)

		err = r.Store(context.Background(), fPath, bytes.NewReader(content))

		require.NoError(t, err)
	})

	t.Run("not enough space", func(t *testing.T) {
		r := New()

		os.SetSpaceLimit(1)

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0750)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())
		testCreateFile(t, fPath, content)

		err = r.Store(context.Background(), fPath, bytes.NewReader(content))

		var errNotEnoughSpace model.NotEnoughSpaceError
		require.ErrorAs(t, err, &errNotEnoughSpace)

		content2, err := io.ReadAll(errNotEnoughSpace.Reader())
		require.NoError(t, err)
		require.Equal(t, content, content2)

		err = errNotEnoughSpace.Close()
		require.NoError(t, err)

		os.SetSpaceLimit(1 << 40)
	})
}
