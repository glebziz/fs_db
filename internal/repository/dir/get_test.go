package dir

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r, ctx := newTestRep(t)

		var (
			dirs = []model.Dir{{
				Id:         gofakeit.UUID(),
				FileCount:  0,
				ParentPath: gofakeit.UUID(),
			}, {
				Id:         gofakeit.UUID(),
				FileCount:  3,
				ParentPath: gofakeit.UUID(),
			}}

			files = []model.ContentFile{{
				Id:         gofakeit.UUID(),
				ParentPath: dirs[1].GetPath(),
			}, {
				Id:         gofakeit.UUID(),
				ParentPath: dirs[1].GetPath(),
			}, {
				Id:         gofakeit.UUID(),
				ParentPath: dirs[1].GetPath(),
			}}
		)

		for _, file := range files {
			testCreateContentFile(ctx, t, r.p, &file)
		}

		for _, dir := range dirs {
			testCreateDir(ctx, t, r.p, &dir)
		}

		actual, err := r.Get(ctx)
		require.NoError(t, err)
		require.Equal(t, dirs, actual)
	})

	t.Run("empty dirs", func(t *testing.T) {
		t.Parallel()

		r, ctx := newTestRep(t)

		actual, err := r.Get(ctx)
		require.NoError(t, err)
		require.Empty(t, actual)
	})
}
