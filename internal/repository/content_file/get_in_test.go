package file

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_GetIn(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		ids    = make([]string, 0, 100)
		files  = make([]model.ContentFile, 0, 100)
		mFiles = make(map[string]model.ContentFile, 100)
	)
	for i := 0; i < 100; i++ {
		file := model.ContentFile{
			Id:     gofakeit.UUID(),
			Parent: gofakeit.UUID(),
		}

		testCreateContentFile(ctx, t, r.p, &file)

		ids = append(ids, file.Id)
		files = append(files, file)
		mFiles[file.Id] = file
	}

	actual, err := r.GetIn(ctx, append(ids, gofakeit.UUID()))
	require.NoError(t, err)

	for _, file := range actual {
		f, ok := mFiles[file.Id]
		require.True(t, ok)
		require.Equal(t, f, file)
	}
}
