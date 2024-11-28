package db

import (
	"bytes"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
)

func TestDb_Tx(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	var (
		c []byte

		key = gofakeit.UUID()
		txs = make([]fs_db.Tx, 3)

		content = [][]byte{
			bytes.Repeat([]byte("0"), 10),
			bytes.Repeat([]byte("1"), 10),
			bytes.Repeat([]byte("2"), 10),
			bytes.Repeat([]byte("3"), 10),
		}
	)

	err := _db.Set(testCtx, key, content[3])
	require.NoError(t, err)

	txs[0], err = _db.Begin(testCtx, fs_db.IsoLevelReadUncommitted)
	require.NoError(t, err)

	// default iso level is read committed
	txs[1], err = _db.Begin(testCtx)
	require.NoError(t, err)

	txs[2], err = _db.Begin(testCtx, fs_db.IsoLevelSerializable)
	require.NoError(t, err)

	// empty txs read contents stored before begin
	for _, _tx := range txs {
		c, err = _tx.Get(testCtx, key)
		require.NoError(t, err)
		require.Equal(t, content[3], c)
	}

	// after storing, read local tx data
	for i, _tx := range txs {
		err = _tx.Set(testCtx, key, content[i])
		require.NoError(t, err)

		c, err = _tx.Get(testCtx, key)
		require.NoError(t, err)
		require.Equal(t, content[i], c)
	}

	// no transaction was committed
	c, err = _db.Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, content[3], c)

	// read uncommitted tx get last stored data
	c, err = txs[0].Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, content[2], c)

	err = txs[0].Commit(testCtx)
	require.NoError(t, err)

	// read committed tx get last committed data
	c, err = txs[1].Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, content[0], c)

	err = txs[1].Commit(testCtx)
	require.NoError(t, err)

	// get last committed data
	c, err = _db.Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, content[1], c)

	// serializable tx get owned data
	c, err = txs[2].Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, content[2], c)

	c, err = txs[2].Get(testCtx, gofakeit.UUID())
	require.ErrorIs(t, err, fs_db.ErrNotFound)
	require.Nil(t, c)

	err = txs[2].Commit(testCtx)
	require.ErrorIs(t, err, fs_db.ErrTxSerialization)

	err = _db.Delete(testCtx, key)
	require.NoError(t, err)

	c, err = _db.Get(testCtx, key)
	require.ErrorIs(t, err, fs_db.ErrNotFound)
	require.Nil(t, c)
}
