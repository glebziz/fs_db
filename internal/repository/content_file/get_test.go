package file

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Get_Success(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		files = []model.ContentFile{{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}, {
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}}
	)

	for _, file := range files {
		testCreateContentFile(ctx, t, r.p, &file)
	}

	for _, file := range files {
		actual, err := r.Get(ctx, file.Id)
		require.NoError(t, err)
		require.Equal(t, &file, actual)
	}
}

func TestRep_Get_Error(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	actual, err := r.Get(ctx, gofakeit.UUID())
	require.ErrorIs(t, err, fs_db.NotFoundErr)
	require.Nil(t, actual)
}
