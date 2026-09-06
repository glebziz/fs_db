package seek_direction

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func TestConvertToWhence(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		dir    store.Seek_Direction
		whence int
	}{
		{
			name:   "SeekStart",
			dir:    store.Seek_SeekStart,
			whence: io.SeekStart,
		},
		{
			name:   "SeekCurrent",
			dir:    store.Seek_SeekCurrent,
			whence: io.SeekCurrent,
		},
		{
			name:   "SeekEnd",
			dir:    store.Seek_SeekEnd,
			whence: io.SeekEnd,
		},
		{
			name:   "unknown",
			dir:    -1,
			whence: io.SeekStart,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			whence := ConvertToWhence(tc.dir)
			require.Equal(t, tc.whence, whence)
		})
	}
}

func TestConvertToGrpc(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		dir  model.SeekDirection
		pb   store.Seek_Direction
	}{
		{
			name: "SeekStart",
			dir:  model.SeekStart,
			pb:   store.Seek_SeekStart,
		},
		{
			name: "SeekCurrent",
			dir:  model.SeekCurrent,
			pb:   store.Seek_SeekCurrent,
		},
		{
			name: "SeekEnd",
			dir:  model.SeekEnd,
			pb:   store.Seek_SeekEnd,
		},
		{
			name: "unknown",
			dir:  -1,
			pb:   store.Seek_SeekStart,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			pb := ConvertToGrpc(tc.dir)
			require.Equal(t, tc.pb, pb)
		})
	}
}
