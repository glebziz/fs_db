package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
)

const (
	testTxId  = "testTxId"
	testTxId2 = "testTxId2"
	testTxId3 = "testTxId3"

	testKey  = "testKey"
	testKey2 = "testKey2"

	testContentId  = "testContentId"
	testContentId2 = "testContentId2"
	testContentId3 = "testContentId3"
	testContentId4 = "testContentId4"
	testContentId5 = "testContentId5"
	testContentId6 = "testContentId6"
)

func requireEqualFiles(t *testing.T, a, b model.File) {
	t.Helper()

	require.Equal(t, a.Key, b.Key)
	require.Equal(t, a.ContentId, b.ContentId)
	require.Equal(t, a.TxId, b.TxId)
}

func (u *useCase) testStore(t *testing.T, f model.File) {
	t.Helper()

	tx, ok := u.txStore.Get(f.TxId)
	if !ok {
		tx = &core.Transaction{}
		u.txStore.Put(f.TxId, tx)
	}

	u.storeToTx(tx, f)
}

func (u *useCase) testAddEmptyTx(t *testing.T, txId string, keys ...string) {
	t.Helper()

	tx := &core.Transaction{}
	u.txStore.Put(txId, tx)
	for _, key := range keys {
		tx.PushBack((&core.Node[model.File]{}).SetV(model.File{
			Key:  key,
			TxId: txId,
		}))
		tx.File(key).PopBack()
	}
}
