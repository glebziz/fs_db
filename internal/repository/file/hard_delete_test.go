package file

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestRep_HardDelete(t *testing.T) {
	t.Parallel()

	r, ctx := newTestRep(t)

	var (
		txId1 = gofakeit.UUID()
		txId2 = gofakeit.UUID()

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
			Key: key2,
		}
		file5 = model.File{
			Key: key1,
		}
		file6 = model.File{
			Key:       key2,
			ContentId: gofakeit.UUID(),
		}
	)

	beforeTs := time.Now().UTC()

	testCreateFile(ctx, t, r.p, txId1, &file1)
	testCreateFile(ctx, t, r.p, txId1, &file2)
	testCreateFile(ctx, t, r.p, txId1, &file3)
	testCreateFile(ctx, t, r.p, txId1, &file4)
	testCreateFile(ctx, t, r.p, txId2, &file3)
	testCreateFile(ctx, t, r.p, txId2, &file4)
	testCreateFile(ctx, t, r.p, txId2, &file5)
	testCreateFile(ctx, t, r.p, txId2, &file6)

	afterTs := time.Now().UTC().Add(time.Hour)

	contentIds, err := r.HardDelete(ctx, txId2, nil)
	require.NoError(t, err)
	require.Equal(t, []string{
		file3.ContentId, file6.ContentId,
	}, contentIds)

	contentIds, err = r.HardDelete(ctx, txId1, &model.FileFilter{
		BeforeTs: ptr.Ptr(beforeTs),
	})
	require.ErrorIs(t, err, fs_db.NotFoundErr)
	require.Empty(t, contentIds)

	contentIds, err = r.HardDelete(ctx, txId1, &model.FileFilter{
		BeforeTs: ptr.Ptr(afterTs),
	})
	require.Equal(t, []string{
		file1.ContentId, file2.ContentId,
	}, contentIds)
}
