package file

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestRep_Store_Success(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		file1 = model.ContentFile{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}

		file2 = model.ContentFile{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}
	)

	err := r.Store(ctx, file1)
	require.NoError(t, err)

	err = r.Store(ctx, file2)
	require.NoError(t, err)

	actual := testGetContentFile(ctx, t, r.p, file1.Id)
	require.Equal(t, &file1, actual)

	actual = testGetContentFile(ctx, t, r.p, file2.Id)
	require.Equal(t, &file2, actual)
}

func TestRep_Create_Error(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		file = model.ContentFile{
			Id:         gofakeit.UUID(),
			ParentPath: gofakeit.UUID(),
		}
	)

	err := r.Store(ctx, model.ContentFile{
		Id: gofakeit.UUID(),
	})
	require.Error(t, err)

	err = r.Store(ctx, file)
	require.NoError(t, err)

	err = r.Store(ctx, model.ContentFile{
		Id:         file.Id,
		ParentPath: gofakeit.UUID(),
	})
	require.Error(t, err)
}
