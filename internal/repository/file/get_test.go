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

func TestRep_Get(t *testing.T) {
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
			Key:       key1,
			ContentId: gofakeit.UUID(),
		}
		file5 = model.File{
			Key:       key2,
			ContentId: gofakeit.UUID(),
		}
		file6 = model.File{
			Key: key2,
		}
	)

	beforeF1Ts := time.Now().UTC()

	testCreateFile(ctx, t, r.p, txId1, &file1)
	testCreateFile(ctx, t, r.p, txId2, &file2)

	afterF2Ts := time.Now().UTC()

	testCreateFile(ctx, t, r.p, txId3, &file3)
	testCreateFile(ctx, t, r.p, txId1, &file4)
	testCreateFile(ctx, t, r.p, txId2, &file5)
	testCreateFile(ctx, t, r.p, txId3, &file6)

	actual, err := r.Get(ctx, txId1, key1, nil)
	require.NoError(t, err)
	require.Equal(t, &file4, actual)

	actual, err = r.Get(ctx, txId1, key2, nil)
	require.ErrorIs(t, err, fs_db.NotFoundErr)
	require.Nil(t, actual)

	actual, err = r.Get(ctx, txId3, key1, &model.FileFilter{
		TxId: ptr.Ptr(txId1),
	})
	require.NoError(t, err)
	require.Equal(t, &file4, actual)

	actual, err = r.Get(ctx, txId2, key1, &model.FileFilter{
		TxId:     ptr.Ptr(txId1),
		BeforeTs: ptr.Ptr(beforeF1Ts),
	})
	require.ErrorIs(t, err, fs_db.NotFoundErr)
	require.Nil(t, actual)

	actual, err = r.Get(ctx, txId2, key1, &model.FileFilter{
		BeforeTs: ptr.Ptr(afterF2Ts),
	})
	require.NoError(t, err)
	require.Equal(t, &file1, actual)

	actual, err = r.Get(ctx, txId3, gofakeit.UUID(), nil)
	require.ErrorIs(t, err, fs_db.NotFoundErr)
	require.Nil(t, actual)
}
