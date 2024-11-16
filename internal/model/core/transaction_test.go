package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func TestTransactions_Get(t *testing.T) {
	for _, tc := range []struct {
		name string
		txId string
		txs  *Transactions
		tx   *Transaction
		ok   bool
	}{
		{
			name: "success",
			txId: testTxId,
			txs: &Transactions{
				store: map[string]*Transaction{
					testTxId: {},
				},
			},
			tx: &Transaction{},
			ok: true,
		},
		{
			name: "nil store",
			txId: testTxId,
			txs:  &Transactions{},
			ok:   false,
		},
		{
			name: "empty store",
			txId: testTxId,
			txs: &Transactions{
				store: map[string]*Transaction{},
			},
			ok: false,
		},
		{
			name: "tx not found",
			txId: "123",
			txs: &Transactions{
				store: map[string]*Transaction{
					testTxId: {},
				},
			},
			ok: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tx, ok := tc.txs.Get(tc.txId)
			require.Equal(t, tc.tx, tx)
			require.Equal(t, tc.ok, ok)
		})
	}
}

func TestTransactions_Put(t *testing.T) {
	for _, tc := range []struct {
		name string
		txs  *Transactions
		tx   *Transaction
	}{
		{
			name: "not nil txs store and not nil tx store",
			txs: &Transactions{
				store: map[string]*Transaction{},
			},
			tx: &Transaction{
				store: map[string]*file{},
			},
		},
		{
			name: "not nil txs store and nil tx store",
			txs: &Transactions{
				store: map[string]*Transaction{},
			},
			tx: &Transaction{},
		},
		{
			name: "nil txs store and not nil tx store",
			txs:  &Transactions{},
			tx: &Transaction{
				store: map[string]*file{},
			},
		},
		{
			name: "nil txs store and nil tx store",
			txs:  &Transactions{},
			tx:   &Transaction{},
		},
		{
			name: "tx already exists",
			txs: &Transactions{
				store: map[string]*Transaction{
					testTxId: {},
				},
			},
			tx: &Transaction{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.txs.Put(testTxId, tc.tx)

			tx, ok := tc.txs.store[testTxId]
			require.True(t, ok)
			require.True(t, tx == tc.tx)
			require.NotNil(t, tx.store)
			require.Empty(t, tx.store)
		})
	}
}

func TestTransactions_Delete(t *testing.T) {
	for _, tc := range []struct {
		name string
		txs  *Transactions
	}{
		{
			name: "success",
			txs: &Transactions{
				store: map[string]*Transaction{
					testTxId: {},
				},
			},
		},
		{
			name: "empty store",
			txs: &Transactions{
				store: map[string]*Transaction{},
			},
		},
		{
			name: "nil store",
			txs:  &Transactions{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			realTx := tc.txs.store[testTxId]

			tx := tc.txs.Delete(testTxId)
			require.Equal(t, realTx, tx)

			tx, ok := tc.txs.store[testTxId]
			require.False(t, ok)
			require.Nil(t, tx)
		})
	}
}

func TestTransaction_Lock(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
	}{
		{
			name: "success",
			tx:   &Transaction{},
		},
		{
			name: "success with nil tx",
			tx:   nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.tx.Lock()
		})
	}
}

func TestTransaction_RLock(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
	}{
		{
			name: "success",
			tx:   &Transaction{},
		},
		{
			name: "success with nil tx",
			tx:   nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.tx.RLock()
		})
	}
}

func TestTransaction_Unlock(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
	}{
		{
			name: "success",
			tx: func() *Transaction {
				tx := &Transaction{}
				tx.Lock()
				return tx
			}(),
		},
		{
			name: "success with nil tx",
			tx: func() *Transaction {
				return nil
			}(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.tx.Unlock()
		})
	}
}

func TestTransaction_RUnlock(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
	}{
		{
			name: "success",
			tx: func() *Transaction {
				tx := &Transaction{}
				tx.RLock()
				return tx
			}(),
		},
		{
			name: "success with nil tx",
			tx: func() *Transaction {
				return nil
			}(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.tx.RUnlock()
		})
	}
}

func TestTransaction_Files(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
	}{
		{
			name: "success",
			tx: &Transaction{
				store: map[string]*file{
					testKey:  {},
					testKey2: {},
				},
			},
		},
		{
			name: "nil tx store",
			tx:   &Transaction{},
		},
		{
			name: "empty tx store",
			tx: &Transaction{
				store: map[string]*file{},
			},
		},
		{
			name: "nil tx",
			tx:   nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			files := tc.tx.Files()

			if tc.tx != nil {
				require.Equal(t, tc.tx.store, files)
			} else {
				require.Nil(t, files)
			}
		})
	}
}

func TestTransaction_File(t *testing.T) {
	for _, tc := range []struct {
		name string
		key  string
		tx   *Transaction
		file *file
	}{
		{
			name: "success",
			key:  testKey,
			tx: &Transaction{
				store: map[string]*file{
					testKey: {},
				},
			},
			file: &file{},
		},
		{
			name: "nil store",
			key:  testKey,
			tx:   &Transaction{},
			file: nil,
		},
		{
			name: "empty store",
			key:  testKey,
			tx: &Transaction{
				store: map[string]*file{},
			},
			file: nil,
		},
		{
			name: "file not found",
			key:  testKey2,
			tx: &Transaction{
				store: map[string]*file{
					testKey: {},
				},
			},
			file: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			file := tc.tx.File(tc.key)
			require.Equal(t, tc.file, file)
		})
	}
}

func TestTransaction_Len(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
		len  int
	}{
		{
			name: "not empty tx store",
			tx: &Transaction{
				store: map[string]*file{
					testKey:  {},
					testKey2: {},
				},
			},
			len: 2,
		},
		{
			name: "empty tx store",
			tx: &Transaction{
				store: map[string]*file{},
			},
			len: 0,
		},
		{
			name: "nil tx store",
			tx:   &Transaction{},
			len:  0,
		},
		{
			name: "nil tx",
			tx:   nil,
			len:  0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.tx.Len()
			require.Equal(t, l, tc.len)
		})
	}
}

func TestTransaction_PushBack(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   func() *Transaction
	}{
		{
			name: "without file",
			tx: func() *Transaction {
				return &Transaction{
					store: map[string]*file{
						testKey2: {},
					},
				}
			},
		},
		{
			name: "with file",
			tx: func() *Transaction {
				f := &file{}
				f.PushBack(&Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId2,
						Seq:       testSeq - 1,
					},
				})

				return &Transaction{
					store: map[string]*file{
						testKey:  f,
						testKey2: {},
					},
				}
			},
		},
		{
			name: "with empty store",
			tx: func() *Transaction {
				return &Transaction{
					store: map[string]*file{},
				}
			},
		},
		{
			name: "with nil store",
			tx: func() *Transaction {
				return &Transaction{}
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			file := model.File{
				Key:       testKey,
				ContentId: testContentId,
				Seq:       testSeq,
			}

			tx := tc.tx()
			tx.PushBack(&Node[model.File]{
				v: file,
			})

			f, ok := tx.store[testKey]
			require.True(t, ok)
			require.Equal(t, file, f.Latest())
		})
	}
}

func TestTransaction_Clear(t *testing.T) {
	for _, tc := range []struct {
		name string
		tx   *Transaction
	}{
		{
			name: "success",
			tx: &Transaction{
				store: map[string]*file{
					testKey:  {},
					testKey2: {},
				},
			},
		},
		{
			name: "success with empty store",
			tx: &Transaction{
				store: map[string]*file{},
			},
		},
		{
			name: "success with nil store",
			tx: &Transaction{
				store: nil,
			},
		},
		{
			name: "success with nil tx",
			tx:   nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.tx.Clear()

			if tc.tx != nil {
				require.Len(t, tc.tx.store, 0)
			}
		})
	}
}
