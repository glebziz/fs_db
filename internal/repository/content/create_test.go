package content

import (
	"context"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	osa "github.com/glebziz/fs_db/internal/adapter/os"
)

func TestRepo_Create(t *testing.T) {
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
					Create(gomock.Any(), filePath).
					Return(nil, nil)
			},
		},
		{
			name: "Create error",
			prepare: func(td *testDeps) {
				td.os.EXPECT().
					Create(gomock.Any(), gomock.Any()).
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
			f, err := r.Create(context.Background(), filePath)

			require.ErrorIs(t, err, tc.err)
			require.Nil(t, f)
		})
	}
}

func TestRepo_Create_Int(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := New(osa.Adapter{})

		f, err := r.Create(context.Background(), path.Join(rootPath, gofakeit.UUID()))
		require.NoError(t, err)
		require.NotNil(t, f)

		f.Close()
	})
}
