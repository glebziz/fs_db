package file

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Delete(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		ids   []string
		files []model.ContentFile
	)
	for i := 0; i < 100; i++ {
		file := model.ContentFile{
			Id:     gofakeit.UUID(),
			Parent: gofakeit.UUID(),
		}

		testCreateContentFile(ctx, t, r.p, &file)

		ids = append(ids, file.Id)
		files = append(files, file)
	}

	err := r.Delete(ctx, append(ids, gofakeit.UUID()))
	require.NoError(t, err)

	for _, id := range ids {
		_, err := r.Get(ctx, id)
		require.ErrorIs(t, err, fs_db.NotFoundErr)
	}
}
