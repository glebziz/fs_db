package content

import (
	"bytes"
	"context"
	"io"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Store(t *testing.T) {
	t.Run("success with existing dir", func(t *testing.T) {
		r := New()

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0666)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())

		err = r.Store(context.Background(), fPath, &model.Content{
			Reader: io.NopCloser(bytes.NewReader(content)),
			Size:   uint64(len(content)),
		})

		require.NoError(t, err)
	})

	t.Run("success with non existing dir", func(t *testing.T) {
		r := New()

		dir := path.Join(rootPath, gofakeit.UUID())
		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())

		err := r.Store(context.Background(), fPath, &model.Content{
			Reader: io.NopCloser(bytes.NewReader(content)),
			Size:   uint64(len(content)),
		})

		require.NoError(t, err)
	})

	t.Run("success with existing file", func(t *testing.T) {
		r := New()

		dir := path.Join(rootPath, gofakeit.UUID())
		err := os.MkdirAll(dir, 0666)
		require.NoError(t, err)

		content := []byte("1234567890")
		fPath := path.Join(dir, gofakeit.UUID())
		testCreateFile(t, fPath, content)

		err = r.Store(context.Background(), fPath, &model.Content{
			Reader: io.NopCloser(bytes.NewReader(content)),
			Size:   uint64(len(content)),
		})

		require.NoError(t, err)
	})
}
