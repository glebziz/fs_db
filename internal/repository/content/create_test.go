package content

import (
	"context"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestRepo_Open(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := New()

		f, err := r.Create(context.Background(), path.Join(rootPath, gofakeit.UUID()))
		require.NoError(t, err)
		require.NotNil(t, f)

		f.Close()
	})
}
