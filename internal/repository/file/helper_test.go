package file

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func newTestRep(t *testing.T) (*rep, context.Context) {
	t.Helper()

	var (
		dbPath = path.Join(os.TempDir(), fmt.Sprintf("test_file_%s.db", gofakeit.UUID()))
	)

	manager, err := db.New(context.Background(), dbPath)
	require.NoError(t, err)
	t.Cleanup(func() {
		manager.Close()
		err = os.Remove(dbPath)
		require.NoError(t, err)
	})

	return New(manager), context.Background()
}

func testCreateFile(ctx context.Context, t *testing.T, p db.Provider, txId string, file *model.File) {
	t.Helper()

	var contentId *string
	if file.ContentId != "" {
		contentId = ptr.Ptr(file.ContentId)
	}

	_, err := p.DB(ctx).Exec(ctx, `
		insert into file(key, content_id, tx_id, ts) VALUES ($1, $2, $3, $4)`,
		file.Key, contentId, txId, time.Now().UTC())
	require.NoError(t, err)
}

func testGetFile(ctx context.Context, t *testing.T, p db.Provider, key string) *model.File {
	t.Helper()

	rows, err := p.DB(ctx).Query(ctx, `
		select key, content_id
		from file where key = $1`, key)
	require.NoError(t, err)
	defer rows.Close()
	require.True(t, rows.Next())

	var (
		file      model.File
		contentId sql.NullString
	)

	err = rows.Scan(&file.Key, &contentId)
	require.NoError(t, err)
	require.NoError(t, rows.Err())

	if contentId.Valid {
		file.ContentId = contentId.String
	}

	return &file
}

func testGetFilesByTx(ctx context.Context, t *testing.T, p db.Provider, txId string) []model.File {
	t.Helper()

	rows, err := p.DB(ctx).Query(ctx, `
		select key, content_id
		from file where tx_id = $1`, txId)
	require.NoError(t, err)
	defer rows.Close()

	var files []model.File
	for rows.Next() {
		var (
			key       string
			contentId sql.NullString
		)

		err = rows.Scan(&key, &contentId)
		require.NoError(t, err)
		require.NoError(t, rows.Err())

		if contentId.Valid {
			files = append(files, model.File{
				Key:       key,
				ContentId: contentId.String,
			})
		}
	}

	return files
}
