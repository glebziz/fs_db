package content

import (
	"context"
	"io"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
)

func TestRep_Get_Success(t *testing.T) {
	r := New()

	dir := path.Join(rootPath, gofakeit.UUID())
	err := os.MkdirAll(dir, 0750)
	require.NoError(t, err)

	content := []byte("1234567890")
	fPath := path.Join(dir, gofakeit.UUID())
	testCreateFile(t, fPath, content)

	c, err := r.Get(context.Background(), fPath)

	require.NoError(t, err)
	require.Equal(t, uint64(len(content)), c.Size)

	actContent, err := io.ReadAll(c.Reader)
	require.NoError(t, err)
	require.Equal(t, content, actContent)

	err = c.Reader.Close()
	require.NoError(t, err)
}

func TestRep_Get_Error(t *testing.T) {
	r := New()

	c, err := r.Get(context.Background(), path.Join(rootPath, gofakeit.UUID()))

	require.ErrorIs(t, err, fs_db.NotFoundErr)
	require.Nil(t, c)
}
