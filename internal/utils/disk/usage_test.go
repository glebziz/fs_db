package disk

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDisk_Usage(t *testing.T) {
	t.Run("success with existing dir", func(t *testing.T) {
		t.Parallel()

		dir, err := os.MkdirTemp("", "test_disk_wrapper")
		require.NoError(t, err)
		t.Cleanup(func() {
			err = os.RemoveAll(dir)
			require.NoError(t, err)
		})

		st, err := GetDisk().Usage(context.Background(), dir)

		require.NoError(t, err)
		require.NotZero(t, st.Total)
		require.NotZero(t, st.Free)
	})

	t.Run("success with non existing dir", func(t *testing.T) {
		t.Parallel()

		dir := path.Join(os.TempDir(), gofakeit.UUID())
		t.Cleanup(func() {
			err := os.RemoveAll(dir)
			require.NoError(t, err)
		})

		st, err := GetDisk().Usage(context.Background(), dir)

		require.NoError(t, err)
		require.NotZero(t, st.Total)
		require.NotZero(t, st.Free)
	})
}
