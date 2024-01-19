package dir

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Create_Success(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		dir1 = model.Dir{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}

		dir2 = model.Dir{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}
	)

	err := r.Create(ctx, dir1)
	require.NoError(t, err)

	err = r.Create(ctx, dir2)
	require.NoError(t, err)

	actual := testGetDir(ctx, t, r.p, dir1.Id)
	require.Equal(t, &dir1, actual)

	actual = testGetDir(ctx, t, r.p, dir2.Id)
	require.Equal(t, &dir2, actual)
}

func TestRep_Create_Error(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	err := r.Create(ctx, model.Dir{
		Id: gofakeit.UUID(),
	})
	require.Error(t, err)
}
