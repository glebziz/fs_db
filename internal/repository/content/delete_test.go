package content

import (
	"context"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/utils/os"
)

func TestRep_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := New()

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
		r := New()

		dir := path.Join(rootPath, gofakeit.UUID())
		fPath := path.Join(dir, gofakeit.UUID())

		err := r.Delete(context.Background(), fPath)
		require.ErrorIs(t, err, fs_db.NotFoundErr)
	})
}
