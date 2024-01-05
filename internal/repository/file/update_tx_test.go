package file

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestRep_UpdateTx(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		txId1 = gofakeit.UUID()
		txId2 = gofakeit.UUID()
		txId3 = gofakeit.UUID()

		key1 = gofakeit.UUID()
		key2 = gofakeit.UUID()

		file1 = model.File{
			Key:       key1,
			ContentId: gofakeit.UUID(),
		}
		file2 = model.File{
			Key:       key2,
			ContentId: gofakeit.UUID(),
		}
		file3 = model.File{
			Key:       key1,
			ContentId: gofakeit.UUID(),
		}
		file4 = model.File{
			Key:       key2,
			ContentId: gofakeit.UUID(),
		}
		file5 = model.File{
			Key:       key1,
			ContentId: gofakeit.UUID(),
		}
		file6 = model.File{
			Key:       key2,
			ContentId: gofakeit.UUID(),
		}
	)

	beforeF1Ts := time.Now().UTC()

	testCreateFile(ctx, t, r.p, txId1, &file1)
	testCreateFile(ctx, t, r.p, txId1, &file2)

	afterF2Ts := time.Now().UTC()

	testCreateFile(ctx, t, r.p, txId2, &file3)
	testCreateFile(ctx, t, r.p, txId2, &file4)
	testCreateFile(ctx, t, r.p, txId3, &file5)
	testCreateFile(ctx, t, r.p, txId3, &file6)

	files := testGetFilesByTx(ctx, t, r.p, txId1)
	require.Equal(t, []model.File{
		file1, file2,
	}, files)

	files = testGetFilesByTx(ctx, t, r.p, txId2)
	require.Equal(t, []model.File{
		file3, file4,
	}, files)

	files = testGetFilesByTx(ctx, t, r.p, txId3)
	require.Equal(t, []model.File{
		file5, file6,
	}, files)

	err := r.UpdateTx(ctx, txId2, txId1, &model.FileFilter{
		BeforeTs: ptr.Ptr(beforeF1Ts),
	})
	require.NoError(t, err)

	files = testGetFilesByTx(ctx, t, r.p, txId2)
	require.Equal(t, []model.File{
		file3, file4,
	}, files)

	err = r.UpdateTx(ctx, txId2, txId1, &model.FileFilter{
		BeforeTs: ptr.Ptr(afterF2Ts),
	})
	require.NoError(t, err)

	files = testGetFilesByTx(ctx, t, r.p, txId2)
	require.Empty(t, files)

	files = testGetFilesByTx(ctx, t, r.p, txId1)
	require.Equal(t, []model.File{
		file1, file2, file3, file4,
	}, files)

	err = r.UpdateTx(ctx, txId3, txId1, nil)
	require.NoError(t, err)

	files = testGetFilesByTx(ctx, t, r.p, txId3)
	require.Empty(t, files)

	files = testGetFilesByTx(ctx, t, r.p, txId1)
	require.Equal(t, []model.File{
		file1, file2, file3, file4, file5, file6,
	}, files)

}
